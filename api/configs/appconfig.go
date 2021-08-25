package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	User     User
	Host     Host
	RabbitMQ RabbitMQ
}

type User struct {
	MongoDb  string
	Database string
}

type Host struct {
	Port int
}

type RabbitMQ struct {
	Host           string
	Port           string
	VHost          string
	Username       string
	Password       string
	ConnectionName string
	Reconnect      struct {
		MaxAttempt int
		Interval   time.Duration
	}
}

func LoadConfig(path, env string) (config Config, err error) {

	filePath := fmt.Sprintf("%s/appsettings.%s.json", path, env)
	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
