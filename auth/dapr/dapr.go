package dapr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	Dapr "github.com/dapr/go-sdk/client"
	Config "github.com/nileshnk/media_upload_poc/auth/config"
)

var DaprClient Dapr.Client


func DaprInit(ctx context.Context) (error) {
	var DaprClientErr error
	addr := Config.GetConfig.Dapr.Host + ":" + strconv.Itoa(Config.GetConfig.Dapr.Port)
	fmt.Println("Dapr Host: ", addr)
	DaprClient, DaprClientErr = Dapr.NewClientWithAddressContext(ctx, addr)
	if DaprClientErr != nil {
		return DaprClientErr
	}
	return nil
}

type EmailPayload struct {
	Recipient []string `json:"recipient"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}


func SendEmail(ctx context.Context, payload EmailPayload) error {
	payloadBytes, payloadBytesErr := json.Marshal(payload)
	if payloadBytesErr != nil {
		return payloadBytesErr
	}

	appID := Config.GetConfig.Dapr.DaprAppCommunication.DaprAppID
	method := Config.GetConfig.Dapr.DaprAppCommunication.Method.SendEmail
	verb := "POST"

	_, err := DaprClient.InvokeMethodWithContent(ctx, appID, method, verb, &Dapr.DataContent{
		ContentType: "application/json",
		Data: payloadBytes,
	})

	if err != nil {
		return err
	}
	return nil
}
