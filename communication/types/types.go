package types

type EmailPayload struct {
	Recipient []string `json:"recipient"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}
