package types

type EmailPayload struct {
	Recipient []string `json:"recipient"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}

type CommonResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}
