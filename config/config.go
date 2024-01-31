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
	// Database is the database configuration
	Database struct {
		// Host is the database host
		Host string
		// Port is the database port
		Port int
		// Name is the database name
		Name string
		// User is the database user
		User string
		// Password is the database password
		Password string
	}
	// Log is the log configuration
	Log struct {
		// Level is the log level
		Level string
		// File is the log file
		File string
	}
	// Auth is the auth configuration
	Auth struct {
		// JWT private key
		PrivateKey string
		// JWT public key
		PublicKey string
	}
	// Email is the email configuration
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
	// Upload is the upload configuration
	Upload struct {
		// Path is the upload path
		Path string
		// MaxSize is the max upload size
		MaxSize int64
	}
}

// C is the global configuration
var GetConfig Config

// Load loads the configuration
func Load() {
	// load from environment variables
	GetConfig.Server.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	GetConfig.Server.Host = os.Getenv("SERVER_HOST")
	GetConfig.Database.Host = os.Getenv("DATABASE_HOST")
	GetConfig.Database.Port, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	GetConfig.Database.Name = os.Getenv("DATABASE_NAME")
	GetConfig.Database.User = os.Getenv("DATABASE_USER")
	GetConfig.Database.Password = os.Getenv("DATABASE_PASSWORD")
	GetConfig.Log.Level = os.Getenv("LOG_LEVEL")
	GetConfig.Log.File = os.Getenv("LOG_FILE")
	GetConfig.Auth.PrivateKey = os.Getenv("JWT_PRIVATE_KEY")
	GetConfig.Auth.PublicKey = os.Getenv("JWT_PUBLIC_KEY")
	GetConfig.Email.Host = os.Getenv("EMAIL_HOST")
	GetConfig.Email.Port, _ = strconv.Atoi(os.Getenv("EMAIL_PORT"))
	GetConfig.Email.User = os.Getenv("EMAIL_USER")
	GetConfig.Email.Password = os.Getenv("EMAIL_PASSWORD")
	GetConfig.Email.From = os.Getenv("EMAIL_FROM")
	GetConfig.Upload.Path = os.Getenv("UPLOAD_PATH")
	GetConfig.Upload.MaxSize, _ = strconv.ParseInt(os.Getenv("UPLOAD_MAX_SIZE"), 10, 64)
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