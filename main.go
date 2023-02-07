package main

import (
	"context"
	"fmt"
	"log"
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
	u4 := GenerateUUID("v4")
	log.Println(formatUUID(u4))
	u1 := GenerateUUID("v1")
	log.Println(formatUUID(u1))
	up := GenerateUUID("pseudo")
	log.Println(formatUUID(up))
}
