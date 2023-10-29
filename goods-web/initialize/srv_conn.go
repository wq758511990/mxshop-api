package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

func InitSrvConn() {
	srvUrl := fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name)
	userConn, err := grpc.Dial(
		srvUrl,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接用户服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
