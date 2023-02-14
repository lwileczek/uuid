package uuid

import (
	"testing"
	"time"
)

func TestV1(t *testing.T) {
	u, err := V1(true)
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

	u2, err := V1(true)
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
		t.Log(u)
		t.Log(u2)
		t.Log("Current Node:", v1State.Node)
		t.Error("Running UUID V1 produces the same UUID each time")
	}

	v1State.Ready = false
	rndu, err := V1(false)
	if err != nil {
		t.Error("Error generating UUID with random bytes instead of MAC address", err)
	}
	if rndu[6]>>4 != 1 {
		t.Error("Version bit not properly set")
		t.Log(u[6] >> 4)
	}

	if !validVariant(rndu) {
		t.Error("The variant bit not properly set")
		t.Log(u[8] >> 4)
	}

	//Verify it is not nill as last check
	for _, b := range u {
		if b != 0 {
			return
		}
	}
	t.Error("The UUID is the nil UUID")
}

func TestSetNode(t *testing.T) {
	t.Parallel()
	us := uuidState{Ready: false}
	if us.Node != [6]byte{0, 0, 0, 0, 0, 0} {
		t.Error("state starting with a non empty Node")
	}
	us.setNode(true)
	if us.Node == [6]byte{0, 0, 0, 0, 0, 0} {
		t.Error("state starting with a non empty Node")
	}
	us.Node = [6]byte{0, 0, 0, 0, 0, 0}
	us.setNode(false)
	if us.Node == [6]byte{0, 0, 0, 0, 0, 0} {
		t.Error("state starting with a non empty Node")
	}
}

//https://www.rfc-editor.org/rfc/rfc4122#section-4.5
func TestRandomMAC(t *testing.T) {
	t.Parallel()
	b := rndNode()
	if b == [6]byte{} {
		t.Error("Random Node doesn't produce a populated array")
	}
	if !(b[0]&1 == 1) {
		t.Log(b[0])
		t.Logf("%b\n", b[0])
		t.Logf("%b\n", b[0]&1)
		t.Error("Least significant bit of the first octet is not 1")
	}

}

func TestRealMAC(t *testing.T) {
	t.Parallel()
	iface, err := defaultHWAddr()
	if err != nil {
		t.Error("error getting the MAC address")
	}
	if len(iface) < 6 {
		t.Error("Iface value is not long enough, likely not MAC")
	}
	for k := 0; k < 6; k++ {
		if iface[k] != byte(0) {
			return
		}
	}
	t.Error("Zero byte array interface")

}

func TestGetTime(t *testing.T) {
	t.Parallel()
	var nulInt uint64
	t0 := getTime()
	time.Sleep(time.Millisecond * 100)
	t1 := getTime()
	if t0 == nulInt {
		t.Error("GetTime produces the nil uint64")
	}

	if t0 == t1 {
		t.Error("Function 'GetTime' always produces the same value")
	}

}

func TestSetClockSeq(t *testing.T) {
	t.Parallel()
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
		_, e := V1(true)
		if e != nil {
			b.Error("Error while benchmarking package", e)
		}
	}
}
