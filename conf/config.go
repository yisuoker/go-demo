package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = new(AppCfg)

type AppCfg struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

func Init(cfgFile string) (err error) {
	viper.SetConfigFile(cfgFile)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(Cfg); err != nil {
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file is changed:", in.Name)
		if err := viper.Unmarshal(Cfg); err != nil {
			return
		}
	})
	return
}
