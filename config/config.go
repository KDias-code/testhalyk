package config

import "github.com/spf13/viper"

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBSSLMode  string
}

func Init() (Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	config := Config{
		ServerPort: viper.GetString("port"),
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBUsername: viper.GetString("db.username"),
		DBPassword: viper.GetString("db.password"),
		DBName:     viper.GetString("db.name"),
		DBSSLMode:  viper.GetString("db.sslmode"),
	}

	return config, nil
}
