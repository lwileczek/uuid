package main

import (
	"crypto/rand"
	"fmt"
)

//Generate UUID v4
func uuidV4() (uuid, error) {
	var rnd [16]byte
	_, err := rand.Read(rnd[:])
	if err != nil {
		fmt.Println("Error: ", err)
		return Nil, err
	}
	rnd[6] = (rnd[6] & 0x0f) | 0x40 // Version 4
	rnd[8] = (rnd[8] & 0x3f) | 0x80 // Variant is 10 - RFC4122
	return rnd, nil
}
