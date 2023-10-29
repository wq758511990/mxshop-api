package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`                    // 手机号码格式 validator
	Scene  string `form:"scene" json:"scene" binding:"required,oneof=register login"` // 发送短信的场景
}
