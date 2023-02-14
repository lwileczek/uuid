package main

import (
	"flag"
	"fmt"

	"github.com/lwileczek/uuid"
)

//No reason other than trying to prevent a runaway process
const (
	maxCount = 5000
)

var (
	uuidTypeFlag string
	createCount  int
	tryMAC       bool
)

func main() {
	flag.StringVar(&uuidTypeFlag, "t", "v1", "The type or version of UUID to generate: v1,v4,Pseudo,null")
	flag.IntVar(&createCount, "n", 1, "How many UUIDs to create")
	flag.BoolVar(&tryMAC, "m", false, "When creating v1 UUIDs, use local MAC addr")
	flag.Parse()

	//Default is to create v1 UUIDs unless other specified
	var uuidType string
	switch uuidTypeFlag {
	case "v1":
		uuidType = uuidTypeFlag
		break
	case "v4":
		uuidType = uuidTypeFlag
		break
	case "pseudo":
		uuidType = uuidTypeFlag
		break
	case "null":
		uuidType = uuidTypeFlag
		break
	default:
		fmt.Println("Did not understand the requested type of UUID")
		fmt.Printf("Got: %s but needs to be one of: v1,v4,Pseudo,null", uuidTypeFlag)
		return
	}

	if (createCount < 1) || (maxCount < createCount) {
		fmt.Printf("Please provide a number 'n' within the interval [1, %d]\n", maxCount)
		return
	}

	for z := 0; z < createCount; z++ {
		u, err := uuid.GenerateUUID(uuidType, tryMAC)
		if err != nil {
			fmt.Printf("There was an error generating the UUID: %s\n", err)
			return
		}
		fmt.Println(uuid.FormatUUID(u))
	}
}
