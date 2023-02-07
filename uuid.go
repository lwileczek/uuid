package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

type uuid [16]byte

var (
	//Nil empty UUID, all Zeros
	Nil uuid
)

// Note - NOT RFC4122 compliant
func pseudoUUID() (uuid, error) {
	var b [16]byte
	//var prng = rand.Reader
	//_, err := io.ReadFull(prng, b[:])
	_, err := rand.Read(b[:])
	if err != nil {
		fmt.Println("Error: ", err)
		return Nil, err
	}
	return b, nil
}

//Convert UUID bytes to a hex string with hyphens
func formatUUID(u uuid) string {
	return fmt.Sprintf("%X-%X-%X-%X-%X", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//GenerateUUID - Generate a UUID
func GenerateUUID(version string) [16]byte {
	switch version {
	case "v4":
		log.Println("v4 requested.")
		generatedUUID, err := uuidV4()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID
	case "v1":
		//log.Println("v1 requested.")
		generatedUUID, err := uuidV1()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID
	case "pseudo":
		log.Println("Pseudo-UUIDrequested.")
		generatedUUID, err := pseudoUUID()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID
	default:
		log.Fatal("Did not understand the version request:", version)
	}
	return Nil
}
