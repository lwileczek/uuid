package main

import (
	"testing"
)

func TestV1(t *testing.T) {
	u, err := uuidV1()
	if err != nil {
		t.Error("Error generating UUID", err)
	}

	if u[6]>>4 != 1 {
		t.Error("Version bit not properly set")
		t.Log(u[6] >> 4)
	}

	if !validVariant(u) {
		t.Error("The variant bit not properly set")
		t.Log(u[8] >> 4)
	}

	u2, err := uuidV1()
	if err != nil {
		t.Error("Error generating Second UUID", err)
	}
	for i, x := range u2[10:] {
		if u[10+i] != x {
			t.Errorf("Node bytes are not the same:\n%s\n%s\n", u, u2)
		}
	}

	sameFlag := true
	for i, x := range u2[0:10] {
		if u[i] != x {
			sameFlag = false
			break
		}
	}
	if sameFlag {
		t.Error("Running UUID V1 produces the same UUID each time")
	}

	//Verify it is not nill as last check
	for _, b := range u {
		if b != 0 {
			return
		}
	}
	t.Error("The UUID is the nil UUID")
}

func BenchmarkV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, e := uuidV1()
		CheckError(e)
	}
}
