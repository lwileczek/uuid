package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lwileczek/uuid"
)

// HandleRequest Handle a request for a new menu description from Lambda
func HandleRequest(ctx context.Context, payload events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request expectedPayload
	UUIDversion := "v4"
	if payload.Body == "" {
		badRequestError := fmt.Errorf("No payload sent")
		response := formatErrorResponse(badRequestError, 400)
		return response, nil
	}
	err := json.Unmarshal([]byte(payload.Body), &request)
	if err != nil {
		log.Println("error unmarshalling request body", err)
		response := formatErrorResponse(err, 500)
		return response, nil
	}

	switch strings.ToLower(request.Version) {
	case "v4":
		break
	case "v1":
		UUIDversion = "v1"
		break
	case "null":
		UUIDversion = "null"
		break
	case "pseudo":
		UUIDversion = "pseudo"
		break
	default:
		badRequestError := fmt.Errorf("Did not understand the version")
		response := formatErrorResponse(badRequestError, 400)
		return response, nil
	}
	var res Response

	if request.Count < 2 {
		u, err := uuid.GenerateUUID(UUIDversion, true)
		if err != nil {
			log.Printf("Error generating v1 UUID: %s\n", err.Error())
		}
		res = Response{
			UUID: uuid.FormatUUID(u),
		}
	} else {
		if request.Count > 5000 {
			request.Count = 5000
		}
		uuids := make([]string, request.Count)
		for i := 0; i < request.Count; i++ {
			u, err := uuid.GenerateUUID(UUIDversion, true)
			if err != nil {
				log.Printf("Error generating v1 UUID: %s\n", err.Error())
			}
			uuids[i] = uuid.FormatUUID(u)
		}
		res = Response{
			UUIDs: uuids,
		}
	}
	stringData, err := json.Marshal(res)
	if err != nil {
		badRequestError := fmt.Errorf("Error creating the JSON response")
		response := formatErrorResponse(badRequestError, 500)
		return response, nil
	}

	response := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(stringData[:]),
	}
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
