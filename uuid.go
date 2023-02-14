package uuid

import (
	"crypto/rand"
	"fmt"
	"log"
)

// CheckError Panic if an error is ever found
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

//FormatUUID Convert UUID bytes to a hex string with hyphens
func FormatUUID(u uuid) string {
	return fmt.Sprintf("%X-%X-%X-%X-%X", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//GenerateUUID - Generate a UUID
func GenerateUUID(version string) ([16]byte, error) {
	switch version {
	case "v4":
		generatedUUID, err := uuidV4()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID, nil
	case "v1":
		generatedUUID, err := uuidV1()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID, nil
	case "pseudo":
		generatedUUID, err := pseudoUUID()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID, nil
	case "null":
		return Nil, nil
	default:
		return Nil, fmt.Errorf("Did not understand the version request: %s", version)
	}
}
