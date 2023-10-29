package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/goods-web/config"
	"mxshop-api/goods-web/proto"
	"strings"
)

var (
	Translator     ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
	GoodsSrvClient proto.GoodsClient
)

type JWTInfo struct {
	SigningKey string
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
