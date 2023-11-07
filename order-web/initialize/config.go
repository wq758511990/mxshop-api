package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/order-web/global"
)

func InitConfig() {
	// 读取本地nacos的配置信息
	v := viper.New()
	debug := v.GetBool("MXSHOP")
	configPath := "config-dev.yaml"
	if debug {
		configPath = "order-web/config-dev.yaml"
	} else {
		configPath = "order-web/config-prd.yaml"
	}
	// 设置路径
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//name := v.Get("name")
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}
	fmt.Println("配置：", global.NacosConfig)
	zap.S().Infof("配置信息: %v", global.NacosConfig)
	// 通过nacos配置信息读取其他配置信息，db等等
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// Create config client for dynamic configuration
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}
	//serverConfig := config.ServerConfig{}
	if err := json.Unmarshal([]byte(content), &global.ServerConfig); err != nil {
		panic(err)
	}

	fmt.Println("globalServer", global.ServerConfig)

	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: global.NacosConfig.DataId,
	//	Group:  global.NacosConfig.Group,
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生了变化...")
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//	},
	//})

}
