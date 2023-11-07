package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"

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
		zap.S().Fatal("[InitSrvConn] 连接商品服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)

	srvOrderUrl := fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.OrderSrvInfo.Name)
	userConn, err = grpc.Dial(
		srvOrderUrl,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接订单服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)

	srvInventoryUrl := fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.InventorySrvInfo.Name)
	userConn, err = grpc.Dial(
		srvInventoryUrl,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接库存服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
