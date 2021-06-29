package setting

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Mode         string
	Port         string
	ReadTimeout  int `mapstructure:"read-timeout"`
	WriteTimeout int `mapstructure:"write-timeout"`
}

type DatabaseConfig struct {
	Type     string
	Username string
	Password string
	Host     string
	Name     string
}

type Setting struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

// Load is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Load(env string) Setting {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on loading configuration file")
	}

	var setting Setting
	err = v.Unmarshal(&setting)
	if err != nil {
		log.Fatal("error on unmarshalling configuration file")
	}

	return setting
}
