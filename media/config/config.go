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
		HTTPPort int
		// Host is the host to listen on
		HTTPHost string

		GRPCPort int
		GRPCHost string
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

	Media struct {
		AWS struct {
			AccessKeyID     string
			SecretAccessKey string
			Region          string
			BucketName      string
			EndPoint        string
		}
	}


}

// C is the global configuration
var GetConfig Config

// Load loads the configuration
func Load() {
	// load from environment variables
	GetConfig.Server.HTTPPort, _ = strconv.Atoi(os.Getenv("SERVER_HTTP_PORT"))
	GetConfig.Server.HTTPHost = os.Getenv("SERVER_HTTP_HOST")
	GetConfig.Server.GRPCPort, _ = strconv.Atoi(os.Getenv("SERVER_GRPC_PORT"))
	GetConfig.Server.GRPCHost = os.Getenv("SERVER_GRPC_HOST")
	GetConfig.Dapr.GRPCPort, _ = strconv.Atoi(os.Getenv("DAPR_GRPC_PORT"))
	GetConfig.Dapr.Host = os.Getenv("DAPR_HOST")
	GetConfig.Media.AWS.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	GetConfig.Media.AWS.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	GetConfig.Media.AWS.Region = os.Getenv("AWS_REGION")
	GetConfig.Media.AWS.BucketName = os.Getenv("AWS_BUCKET_NAME")
	GetConfig.Media.AWS.EndPoint = os.Getenv("AWS_ENDPOINT")
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
