package communication

import (
	"fmt"
	"net/smtp"
	"strconv"

	Config "github.com/nileshnk/media_upload_poc/config"
)

func SendEmail(recipient []string, subject string, body string) {
	// Message.
	message := []byte("This is a test email message.")
	
	// Authentication.

	smtpHost := Config.GetConfig.Email.Host
	smtpPort := strconv.Itoa(Config.GetConfig.Email.Port)
	smtpPassword := Config.GetConfig.Email.Password
	smtpUsername := Config.GetConfig.Email.User

	validateConfig := Config.Validate([]string{smtpHost, smtpPort, smtpPassword, smtpUsername})
	if !validateConfig {
		fmt.Println("Invalid Config")
		return 
	}


	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUsername, recipient, message)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	
	fmt.Println("Email Sent Successfully!")
	
}