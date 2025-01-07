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
	Dev_BypassOtp              bool `mapstructure:"DEVMODE_BYPASS_OTP"`
	Dev_AllowSendingEmails     bool `mapstructure:"DEVMODE_ALLOW_SENDING_EMAILS"`
	Dev_LogCredentials         bool `mapstructure:"DEVMODE_LOG_CREDENTIALS"`
	Dev_Mode                   bool `mapstructure:"DEV_MODE"`
}

type Twilio struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
}

type Emailing struct {
	FromEmail         string `mapstructure:"EMAIL_FROM"`
	AppPassword       string `mapstructure:"EMAIL_APP_PASSWORD"`
	SmtpServerAddress string `mapstructure:"SMTP_SERVER_ADDRESS"`
	SmtpsPort         string `mapstructure:"SMTPS_PORT"`
}

var (
	InitialData struct {
		SuperAdminUsername string `mapstructure:"INITIAL_SUPER_ADMIN_USERNAME"`
		SuperAdminPassword string `mapstructure:"INITIAL_SUPER_ADMIN_PASSWORD"`
	}

	Configs struct {
		PostgresConn      `mapstructure:",squash"`
		Env               `mapstructure:",squash"`
		DevelopmentConfig `mapstructure:",squash"`
		Twilio            `mapstructure:",squash"`
		Emailing          `mapstructure:",squash"`
	}
)

func loadConfig() {
	viper.AutomaticEnv()
	if environment := viper.GetString("ENVIRONMENT"); environment == "LOCAL" || environment == "" {

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

		err = viper.Unmarshal(&InitialData)
		if err != nil {
			log.Fatalln("error occured while writing env values onto variables, error:", err)
		}
	} else {
		Configs.PostgresConn.DbHost = viper.GetString("DB_HOST")
		Configs.PostgresConn.DbUser = viper.GetString("DB_USER")
		Configs.PostgresConn.DbPassword = viper.GetString("DB_PASSWORD")
		Configs.PostgresConn.DbPort = viper.GetString("DB_PORT")
		Configs.PostgresConn.DbName = viper.GetString("DB_NAME")
		Configs.PostgresConn.DbSslMode = viper.GetString("DB_SSL_MODE")

		Configs.Env.Port = viper.GetString("PORT")
		Configs.Env.JwtSecretKey = viper.GetString("JWT_SECRET_KEY")
		Configs.Env.Environment = viper.GetString("ENVIRONMENT")
		Configs.Env.CORSAllowedOrigins = viper.GetString("CORS_ALLOWED_ORIGINS")

		Configs.DevelopmentConfig.Dev_AllowUniversalPassword = viper.GetBool("DEVMODE_ALLOW_UNIVERSAL_PASSWORD")
		Configs.DevelopmentConfig.Dev_BypassOtp = viper.GetBool("DEVMODE_BYPASS_OTP")
		Configs.DevelopmentConfig.Dev_AllowSendingEmails = viper.GetBool("DEVMODE_ALLOW_SENDING_EMAILS")
		Configs.DevelopmentConfig.Dev_LogCredentials = viper.GetBool("DEVMODE_LOG_CREDENTIALS")
		Configs.DevelopmentConfig.Dev_Mode = viper.GetBool("DEV_MODE")

		Configs.Twilio.AccountSid = viper.GetString("TWILIO_ACCOUNT_SID")
		Configs.Twilio.AuthToken = viper.GetString("TWILIO_AUTH_TOKEN")
		Configs.Twilio.ServiceSid = viper.GetString("TWILIO_SERVICE_SID")

		Configs.Emailing.FromEmail = viper.GetString("EMAIL_FROM")
		Configs.Emailing.AppPassword = viper.GetString("EMAIL_APP_PASSWORD")
		Configs.Emailing.SmtpServerAddress = viper.GetString("SMTP_SERVER_ADDRESS")
		Configs.Emailing.SmtpsPort = viper.GetString("SMTPS_PORT")

		InitialData.SuperAdminUsername = viper.GetString("INITIAL_SUPER_ADMIN_USERNAME")
		InitialData.SuperAdminPassword = viper.GetString("INITIAL_SUPER_ADMIN_PASSWORD")
	}

	fmt.Println("Envirnment variables loaded successfully")
}
