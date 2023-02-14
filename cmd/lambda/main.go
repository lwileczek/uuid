package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lwileczek/uuid"
)

// CheckError Panic if an error is ever found
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//HandleRequest Handle a request and return data back
func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fmt.Printf("%v\n", event)
	return "Hello from Lambda!", nil
}

func main() {
	//lambda.Start(HandleRequest)
	u4, err := uuid.GenerateUUID("v4")
	if err != nil {
		log.Printf("Error generating v4 UUID: %s\n", err.Error())
	}
	log.Println(uuid.FormatUUID(u4))
	up, err := uuid.GenerateUUID("pseudo")
	if err != nil {
		log.Printf("Error generating pseudo UUID: %s\n", err.Error())
	}
	log.Println(uuid.FormatUUID(up))
	log.Println("10 v1 UUIDs")
	for i := 0; i < 10; i++ {
		u1, err := uuid.GenerateUUID("v1")
		if err != nil {
			log.Printf("Error generating v1 UUID: %s\n", err.Error())
		}
		log.Println(uuid.FormatUUID(u1))
	}
}
