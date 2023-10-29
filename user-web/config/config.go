package config

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	Host        string        `mapstructure:"host" json:"host"`
	Tags        []string      `mapstructure:"tags" json:"tags"`
	UserSrvInfo UserSrvConfig `mapstructure:"user-srv" json:"user-srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}

type AliSmsConfig struct {
	ApiKey     string `mapstructure:"key" json:"key"`
	ApiSecrect string `mapstructure:"secrect" json:"secrect"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
