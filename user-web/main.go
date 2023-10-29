package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop-api/goods-web/utils/register/consul"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	"mxshop-api/user-web/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, _ := zap.NewDevelopment()
	// 初始化配置文件
	initialize.InitConfig()
	//初始化翻译器
	if err := initialize.InitValidator("zh"); err != nil {
		fmt.Printf("初始化翻译器错误, err = %s", err.Error())
		return
	}
	zap.ReplaceGlobals(logger)
	// 初始化路由
	router := initialize.Routers()
	// 初始化srv服务连接
	initialize.InitSrvConn()
	port, _ := utils.GetFreePort()
	//if err == nil {
	//	global.ServerConfig.Port = port
	//}
	// 注册consul服务
	registerClient := consul.NewRegister(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		panic("服务注册失败")
	}

	zap.S().Info("port: ", port)
	zap.S().Infof("启动服务器，端口 %d", global.ServerConfig.Port)
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	// 接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = registerClient.DeRegister(serviceId)
	if err != nil {
		zap.S().Panic("注销失败：", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
