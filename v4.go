package uuid

import (
	"crypto/rand"
	"fmt"
)

//Generate UUID v4
func uuidV4() (uuid, error) {
	_, err := rand.Read(blankUUID[:])
	if err != nil {
		fmt.Println("Error: ", err)
		return Nil, err
	}
	blankUUID[6] = (blankUUID[6] & 0x0f) | 0x40 // Version 4
	blankUUID[8] = (blankUUID[8] & 0x3f) | 0x80 // Variant is 10 - RFC4122
	return blankUUID, nil
}
