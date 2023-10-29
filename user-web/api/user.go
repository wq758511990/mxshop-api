package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"
	"net/http"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换为http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
				break
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
				break
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
				break
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
		}
		return
	}
}

func HandleValidatorError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": global.RemoveTopStruct(errs.Translate(global.Translator)),
	})
}

func GetUserList(ctx *gin.Context) {
	// 从注册中心获取到用户服务的信息
	// 调用接口
	pn := ctx.DefaultQuery("pn", "0")
	psize := ctx.DefaultQuery("psize", "10")
	pnInt, _ := strconv.Atoi(pn)
	pSize, _ := strconv.Atoi(psize)

	list, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSize),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询列表失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)
	result := make([]interface{}, 0)
	for _, value := range list.Data {
		rsp := response.UserResponse{
			Id:       value.Id,
			Nickname: value.NickName,
			Birthday: time.Time(time.Unix(int64(value.Birthday), 0)).Format("2006-01-01"),
			//Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender: value.Gender,
			Mobile: value.Mobile,
		}

		result = append(result, rsp)
	}
	ctx.JSON(http.StatusOK, result)

	zap.S().Debug("获取用户列表页")
}

func PasswordLogin(ctx *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		// 返回错误信
		HandleValidatorError(ctx, err)
		return
	}
	// 验证码做验证
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
	// 登录
	resp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	}
	// 查询到用户，未检查密码，需要检查
	if passResp, _ := global.UserSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: resp.Password,
	}); passResp.Success != true {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"password": "密码错误",
		})
	} else {
		j := middlewares.NewJWT()
		claims := models.CustomClaims{
			ID:          uint(resp.Id),
			NickName:    resp.NickName,
			AuthorityId: uint(resp.Role),
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),               // 签名生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
				Issuer:    "imooc",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生成token失败",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id":        resp.Id,
			"nickname":  resp.NickName,
			"token":     token,
			"expire_at": (time.Now().Unix() + 60*60*24*30) * 1000,
		})
	}
}

func Register(ctx *gin.Context) {
	// 用户注册
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	// 验证
	addr := fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if err != nil || value != registerForm.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	// 连接用户grpc
	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		NickName: registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("[Register] 【新建用户失败】: %s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"nickname": user.NickName,
	})
}
