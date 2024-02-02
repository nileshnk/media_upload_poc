package config

import (
	"errors"
	"os"
	"strconv"
)

// Config is the configuration structure
type Config struct {
	// Server is the server configuration
	Server struct {
		// Port is the port to listen on
		Port int
		// Host is the host to listen on
		Host string
	}

	Email struct {
		// Host is the email host
		Host string
		// Port is the email port
		Port int
		// User is the email user
		User string
		// Password is the email password
		Password string
		// From is the email from
		From string
	}

	Dapr struct {
		// Port is the port to listen on
		GRPCPort int
		// Host is the host to listen on
		Host string

		DaprAppID string

		DaprAppCommunication struct {
			DaprAppID string
			Method    struct {
				SendEmail string
			}
		}
	}
}

// C is the global configuration
var GetConfig Config

// Load loads the configuration
func Load() {
	// load from environment variables
	GetConfig.Server.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	GetConfig.Server.Host = os.Getenv("SERVER_HOST")
	GetConfig.Dapr.GRPCPort, _ = strconv.Atoi(os.Getenv("DAPR_GRPC_PORT"))
	GetConfig.Dapr.Host = os.Getenv("DAPR_HOST")
	GetConfig.Dapr.DaprAppID = "auth"
	GetConfig.Dapr.DaprAppCommunication.DaprAppID = "communication"
	GetConfig.Dapr.DaprAppCommunication.Method.SendEmail = "email_delivery"
	GetConfig.Email.Host = os.Getenv("EMAIL_HOST")
	GetConfig.Email.Port, _ = strconv.Atoi(os.Getenv("EMAIL_PORT"))
	GetConfig.Email.User = os.Getenv("EMAIL_USER")
	GetConfig.Email.Password = os.Getenv("EMAIL_PASSWORD")
	GetConfig.Email.From = os.Getenv("EMAIL_FROM")

}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func Validate(strArr []string) error {
	for _, str := range strArr {
		if IsEmpty(str) {
			return errors.New("Empty string")
		}
	}
	return nil
}
