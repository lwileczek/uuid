package uuid

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

func TestGetTime(t *testing.T) {
	var nulInt uint64
	t0 := getTime()
	t.Log("First time uint64:", t0)
	t1 := getTime()
	t.Log("Secod time uint64:", t1)

	if t0 == nulInt {
		t.Error("GetTime produces the nil uint64")
	}

	if t0 == t1 {
		t.Error("Get time always produces the same value")
	}

}

func TestSetClockSeq(t *testing.T) {
	us := uuidState{Ready: false}
	err := us.setClockSeq()
	if err != nil {
		t.Error("Error setting the clock sequence", err)
	}
	var emptyInt uint16
	if us.ClockSeq == emptyInt {
		t.Error("Clock sequence set to zero")
	}
	firstClockSeq := us.ClockSeq
	t.Log("First Clock Sequence", firstClockSeq)
	err = us.setClockSeq()
	if err != nil {
		t.Error("Error setting the clock sequence the second time", err)
	}
	if firstClockSeq == us.ClockSeq {
		t.Error("Clock sequence always returns the same thing", us.ClockSeq)
	}
}

func BenchmarkV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, e := uuidV1()
		CheckError(e)
	}
}
