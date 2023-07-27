package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = new(AppCfg)

func Init(cfgFile string) (err error) {
	if "" == cfgFile {
		cfgFile = "./conf/config.yaml"
	}
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
