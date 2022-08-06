package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Database struct {
	DBHost            string `mapstructure:"db_host"`
	DBPort            string `mapstructure:"db_port"`
	DBName            string `mapstructure:"db_name"`
	DBUsername        string `mapstructure:"db_username"`
	DBPassword        string `mapstructure:"db_password"`
	DBPostgresSslMode string `mapstructure:"db_postgres_ssl_mode"`
}

type Configuration struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file: ", err)
		os.Exit(1)
	}
}

func InitConfigDsn() string {
	initConfig()

	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("[CONFIG][InitConfigDsn] Unable to decode into struct:", err)
	}

	dsn := `postgres://` + config.DBUsername + `:` + config.DBPassword + `@` + config.DBHost + `:` + config.DBPort + `/` + config.DBName + `?sslmode=` + config.DBPostgresSslMode

	return dsn
}

func InitConfigServer() string {
	initConfig()

	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("[CONFIG][InitConfigServer] Uncable to decode into struct:", err)
	}

	port := `:` + fmt.Sprint(config.Port)

	return port
}
