package conf

type AppCfg struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`

	*DbCfg    `mapstructure:"mysql"`
	*RedisCfg `mapstructure:"redis"`
}
