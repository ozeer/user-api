package initialize

import (
	"fmt"
	"user-api/config"
	"user-api/global"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.AutomaticEnv()
	return viper.GetString(key)
}

func InitConfig() {
	v := viper.New()

	env := GetEnv("ENV")

	if env == "" {
		env = global.ENV_DEV
	} else {
		if env != global.ENV_DEV && env != global.ENV_PROD {
			env = global.ENV_DEV
		}
	}
	configFileName := fmt.Sprintf("config-%s.yaml", env)

	v.SetConfigName(configFileName) // name of config file (without extension)
	v.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(".")            // optionally look for config in the working directory
	err := v.ReadInConfig()         // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	conf := config.Conf{}
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	global.Conf = &conf
}
