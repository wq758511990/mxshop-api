package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

func InitSrvConn() {
	srvUrl := fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name)
	conn, err := grpc.Dial(
		srvUrl,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接商品服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(conn)

	srvUserOpUrl := fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserOpSrvInfo.Name)
	conn, err = grpc.Dial(
		srvUserOpUrl,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接订单服务失败")
	}
	global.UserFavClient = proto.NewUserFavClient(conn)
	global.MessageClient = proto.NewMessageClient(conn)
	global.AddressClient = proto.NewAddressClient(conn)
}
