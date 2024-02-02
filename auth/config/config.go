package config

import (
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

	Dapr struct {
		// Port is the port to listen on
		GRPCPort int
		// Host is the host to listen on
		Host string

		DaprAppID string

		DaprAppCommunication struct {
			DaprAppID      string
			DaprPubSubName string
			TopicName      struct {
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
	GetConfig.Dapr.DaprAppCommunication.TopicName.SendEmail = "email_delivery"

}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func Validate(strArr []string) bool {
	for _, str := range strArr {
		if IsEmpty(str) {
			return false
		}
	}
	return true
}
