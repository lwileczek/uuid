package uuid

import (
	"crypto/rand"
	"fmt"
	"log"
)

type uuid [16]byte

var (
	//Nil empty UUID, all Zeros
	Nil       uuid
	blankUUID uuid
)

// Note - NOT RFC4122 compliant
func pseudoUUID() (uuid, error) {
	_, err := rand.Read(blankUUID[:])
	if err != nil {
		fmt.Println("Error: ", err)
		return Nil, err
	}
	return blankUUID, nil
}

//FormatUUID Convert UUID bytes to a hex string with hyphens
func FormatUUID(u uuid) string {
	return fmt.Sprintf("%X-%X-%X-%X-%X", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//GenerateUUID - Generate a UUID
// Params:
//	version: The version of UUID to create, e.g., v1,v4,v5, etc.
//  tryMAC: If using v1, try checking for a MAC addr or create random
//	Some people may not want to use their MAC address for security reasons
func GenerateUUID(version string, tryMAC bool) ([16]byte, error) {
	switch version {
	case "v4":
		generatedUUID, err := uuidV4()
		if err != nil {
			log.Fatal(err)
		}
		return generatedUUID, nil
	case "v1":
		generatedUUID, err := V1(tryMAC)
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
