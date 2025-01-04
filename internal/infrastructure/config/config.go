package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

type Env struct {
	Port               string `mapstructure:"PORT"`
	JwtSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	Environment        string `mapstructure:"ENVIRONMENT"`
	CORSAllowedOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
}

type PostgresConn struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbSslMode  string `mapstructure:"DB_SSL_MODE"`
}

type DevelopmentConfig struct {
	Dev_InitDbEmpty            bool
	Dev_AllowUniversalPassword bool
	Dev_AutoMigrateDbOnStart   bool
	Dev_AllowSystemEndpoints   bool
}

var (
	InitialData struct {
		SuperAdminUsername string
		SuperAdminPassword string
	}

	Configs struct {
		PostgresConn      `mapstructure:",squash"`
		Env               `mapstructure:",squash"`
		DevelopmentConfig `mapstructure:",squash"`
	}
)

func loadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error occured while reading env variables, error:", err)
	}

	err = viper.Unmarshal(&Configs)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	InitialData.SuperAdminUsername = viper.GetString("SUPER_ADMIN_USERNAME")
	InitialData.SuperAdminPassword = viper.GetString("SUPER_ADMIN_PASSWORD")

	fmt.Println("Envirnment variables loaded successfully")
}
