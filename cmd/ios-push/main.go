package main

import (
	"log"
	"os"
	"time"

	"github.com/open-libs/tpns-api-go-sdk/pkg/client"
	"github.com/open-libs/tpns-api-go-sdk/pkg/client/endpoints"
	"github.com/open-libs/tpns-api-go-sdk/pkg/models"
)

func main() {

	accessID := os.Args[1]
	secretKey := os.Args[2]
	token := os.Args[3]

	customContent := ""
	if len(os.Args) > 4 {
		customContent = os.Args[4]
	}

	env := "dev"
	if len(os.Args) > 5 {
		env = os.Args[5]
	}
	c := &client.Client{}
	c.Init(endpoints.Guangzhou).WithSecretId(accessID, secretKey)
	now := time.Now()

	req := &models.IOSPushRequest{
		BaseRequest:  &client.BaseRequest{Path: "/v3/push/app"},
		AudienceType: models.AudienceToken,
		TokenList:    []string{token},
		MessageType:  models.MessageTypeNotify,
		Message: models.IOSMessage{
			Title:   "test",
			Content: "This is Content @" + now.Format(time.RFC3339),
			IOS:     models.IOSContent{CustomContent: customContent},
		},
		Environment: env, // product or dev
	}
	log.Printf("AccessID: %s\n", accessID)
	log.Printf("IOS Request: %s\n", req.ToJsonString())

	resp := &client.TPNSBaseResponse{}

	if err := c.Send(req, resp); err != nil {
		log.Printf("Error: %s\n", err.Error())
	} else {
		log.Printf("IOS Response: %s\n", resp.ToJsonString())
	}
}
