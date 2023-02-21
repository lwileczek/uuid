package main

import "github.com/aws/aws-lambda-go/events"

type expectedPayload struct {
	Count   int    `json:"count"`
	Version string `json:"version"`
	Word    string `json:"word,omitempty"`
}

// Response Simple wrapper for JOSN payload
type Response struct {
	UUID  string   `json:"uuid,omitempty"`
	UUIDs []string `json:"uuids,omitempty"`
}

func formatErrorResponse(e error, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      code,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":     "text/plain",
			"x-amzn-ErrorType": "502",
		},
		Body: e.Error(),
	}
}
