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
	Dev_AllowUniversalPassword bool `mapstructure:"DEVMODE_ALLOW_UNIVERSAL_PASSWORD"`
	Dev_AutoMigrateDbOnStart   bool `mapstructure:"DEVMODE_AUTO_MIGRATE_DB_ON_START"`
	Dev_BypassOtp              bool `mapstructure:"DEVMODE_BYPASS_OTP"`
	Dev_Mode                   bool `mapstructure:"DEV_MODE"`
}

type Twilio struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
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
		Twilio            `mapstructure:",squash"`
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

	InitialData.SuperAdminUsername = viper.GetString("INITIAL_SUPER_ADMIN_USERNAME")
	InitialData.SuperAdminPassword = viper.GetString("INITIAL_SUPER_ADMIN_PASSWORD")

	fmt.Println("Envirnment variables loaded successfully")
}
