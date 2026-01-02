package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
	StriprSecret          string
	SuccessUrl            string
	CancelUrl             string
}

func SetupEnv() (cfg AppConfig, err error) {

	//	if os.Getenv("APP_ENV") == "dev" {
	godotenv.Load()
	//	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New(" httpport var not found")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New(" DSN var not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New(" app secret var not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: dsn, AppSecret: appSecret,
		TwilioAccountSid:      os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:       os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhoneNumber: os.Getenv("TWILIP_FROM_PHONE_NUMBER"),
		StriprSecret:          os.Getenv("STRIPE_SECRET"),
		SuccessUrl:            os.Getenv("SUCCESS_URL"),
		CancelUrl:             os.Getenv("CANCEL_URL"),
	}, nil
}
