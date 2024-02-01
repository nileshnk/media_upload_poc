package dapr

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	Dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprCommon "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	Config "github.com/nileshnk/media_upload_poc/communication/config"
)

var DaprClient Dapr.Client
var DaprServer daprCommon.Service

func DaprInit(ctx context.Context) (error) {
	var DaprClientErr error

	addr := Config.GetConfig.Dapr.Host + ":" + strconv.Itoa(Config.GetConfig.Dapr.Port)
	Config.Validate([]string{addr})

	DaprClient, DaprClientErr = Dapr.NewClientWithAddressContext(ctx, addr)
	if DaprClientErr != nil {
		return DaprClientErr
	}
	
	DaprServer = daprd.NewService(addr)

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

func DaprServiceSetup (ctx context.Context) {
	// Add a handler to the service server
	emailSubscription := &common.Subscription{
		PubsubName: "communication",
		Topic:      "email_delivery",
		Route:      "/send-email",
	}

	if err := DaprServer.AddTopicEventHandler(emailSubscription, emailSubscriptionHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := DaprServer.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}

}

func emailSubscriptionHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	var payload EmailPayload

	jsonParseErr := json.Unmarshal([]byte(*e.Data.(*string)), &payload)

	if jsonParseErr != nil {
		return true, jsonParseErr
	}

	sendEmailErr := SendEmail(ctx, payload)
	if sendEmailErr != nil {
		return true, sendEmailErr
	}

	return false, nil
}
