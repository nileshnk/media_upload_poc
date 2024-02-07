package dapr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	Dapr "github.com/dapr/go-sdk/client"
	Config "github.com/nileshnk/media_upload_poc/media/config"
	Types "github.com/nileshnk/media_upload_poc/media/types"
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

type CommonResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func ServiceInvoke(ctx context.Context, serviceName string, methodName string, payload interface{}) ([]byte, error) {

	resp, err := DaprClient.InvokeMethodWithCustomContent(ctx, serviceName, methodName, "POST", "application/json", payload)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func InvokeAuthTokenValidate(ctx context.Context, token string) bool {
	jsonToken := map[string]string{"token": token}
	resp, err := ServiceInvoke(ctx, "auth", "/validate_token", jsonToken)
	if err != nil {
		fmt.Println(err)
	}
	jsonResp := CommonResponse{}
	json.Unmarshal(resp, &jsonResp)

	if jsonResp.Success == false {
		return false
	}
	fmt.Println(string(resp))
	fmt.Println(jsonResp)
	return true
}

func TestServiceInvoke(ctx context.Context) {
	resp, err := ServiceInvoke(ctx, "communication", "test", []byte("hello from media"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
}
