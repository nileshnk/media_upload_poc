package types

import "github.com/google/uuid"

type EmailPayload struct {
	Recipient []string `json:"recipient"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}

type TokenPayload struct {
	UserId uuid.UUID `json:"user_id"`
}

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CommonResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}
