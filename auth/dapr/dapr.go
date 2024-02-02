package dapr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	Dapr "github.com/dapr/go-sdk/client"
	Config "github.com/nileshnk/media_upload_poc/auth/config"
	Types "github.com/nileshnk/media_upload_poc/auth/types"
)

var DaprClient Dapr.Client

func DaprInit(ctx context.Context) error {
	var DaprClientErr error
	addr := Config.GetConfig.Dapr.Host + ":" + strconv.Itoa(Config.GetConfig.Dapr.GRPCPort)
	fmt.Println("Dapr Host: ", addr)
	DaprClient, DaprClientErr = Dapr.NewClientWithAddressContext(ctx, addr)
	if DaprClientErr != nil {
		return DaprClientErr
	}
	return nil
}

func SendEmail(ctx context.Context, payload Types.EmailPayload) error {
	payloadBytes, payloadBytesErr := json.Marshal(payload)
	if payloadBytesErr != nil {
		return payloadBytesErr
	}

	pubsubName := Config.GetConfig.Dapr.DaprAppCommunication.DaprAppID
	topicName := Config.GetConfig.Dapr.DaprAppCommunication.TopicName.SendEmail

	err := DaprClient.PublishEvent(ctx, pubsubName, topicName, payloadBytes)

	if err != nil {
		return err
	}
	return nil
}
