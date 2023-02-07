package main

import (
	"testing"
)

func TestV4(t *testing.T) {
	u, err := uuidV4()
	if err != nil {
		t.Error("Error generating UUID", err)
	}

	if u[6]>>4 != 4 {
		t.Error("Version bit not properly set")
		t.Log(u[6] >> 4)
	}

	if !validVariant(u) {
		t.Error("The variant bit not properly set")
		t.Log(u[8] >> 4)
	}

	for _, b := range u {
		if b != 0 {
			return
		}
	}
	t.Error("The UUID is the nil UUID")
}

func validVariant(uid uuid) bool {
	switch {
	case (uid[8] & 0xc0) == 0x80:
		return true
	case (uid[8] & 0xe0) == 0xc0:
		return true
	case (uid[8] & 0xe0) == 0xe0:
		return true
	default:
		return false
	}
}

func BenchmarkV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidV4()
	}
}
