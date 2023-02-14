package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

/* From RFC4122
 * Field                  Data Type     Octet  Note #
 *
 * time_low               unsigned 32   0-3    The low field of the
 *                        bit integer          timestamp
 *
 * time_mid               unsigned 16   4-5    The middle field of the
 *                        bit integer          timestamp
 *
 * time_hi_and_version    unsigned 16   6-7    The high field of the
 *                        bit integer          timestamp multiplexed
 *                                             with the version number
 *
 * clock_seq_hi_and_rese  unsigned 8    8      The high field of the
 * rved                   bit integer          clock sequence
 *                                             multiplexed with the
 *                                             variant
 *
 * clock_seq_low          unsigned 8    9      The low field of the
 *                        bit integer          clock sequence
 *
 * node                   unsigned 48   10-15  The spatially unique
 *                        bit integer          node identifier
 */

const (
	lillian    = 2299160          // Julian day of 15 Oct 1582
	unix       = 2440587          // Julian day of 1 Jan 1970
	epoch      = unix - lillian   // Days between epochs
	g1582      = epoch * 86400    // seconds between epochs
	g1582ns100 = g1582 * 10000000 // 100s of a nanoseconds between epochs
)

//v1State State struct to manage timestamp state when generating new UUIDs
var v1State = uuidState{Ready: false}

//Generate UUID v1
func uuidV1() (uuid, error) {
	var uuid [16]byte
	//Another option is to use sync.Once
	if v1State.Ready {
		//increment and continue
		tm := getTime()
		if tm <= v1State.Timestamp {
			v1State.ClockSeq++
		} else {
			v1State.Timestamp = tm
		}
	} else {
		// construct the state
		v1State.setNode()
		err := v1State.setClockSeq()
		if err != nil {
			log.Printf("Error setting the clock sequence\n%s", err.Error())
		}
		v1State.Timestamp = getTime()
		v1State.Ready = true
	}
	ut := uuidTime{}
	ut.TimeLow = uint32(v1State.Timestamp & 0xFFFFFFFF)
	ut.TimeMid = uint16((v1State.Timestamp >> 32) & 0xFFFF)
	ut.TimeHiAndVersion = uint16((v1State.Timestamp >> 48) & 0x0FFF)
	ut.TimeHiAndVersion |= (1 << 12)
	ut.ClockSeqLow = uint8(v1State.ClockSeq & 0xFF)
	ut.ClockSeqHi = uint8((v1State.ClockSeq & 0x3F00) >> 8)
	ut.ClockSeqHi |= 0x80

	//Now assign all the bits to our UUID [16]byte
	binary.BigEndian.PutUint32(uuid[0:4], ut.TimeLow)
	binary.BigEndian.PutUint16(uuid[4:6], ut.TimeMid)
	binary.BigEndian.PutUint16(uuid[6:8], ut.TimeHiAndVersion)
	uuid[8] = ut.ClockSeqHi
	uuid[9] = ut.ClockSeqLow
	for i, b := range v1State.Node {
		uuid[10+i] = b
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x10 // Version 1
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10 - RFC4122

	return uuid, nil
}

type uuidState struct {
	Ready     bool
	ClockSeq  uint16 //uint8 Originally had uint8, is that from the time below?
	Timestamp uint64
	Node      [6]byte
}
type uuidTime struct {
	TimeLow          uint32
	TimeMid          uint16
	TimeHiAndVersion uint16
	ClockSeqHi       uint8
	ClockSeqLow      uint8
}

func getTime() uint64 {
	meow := g1582ns100 + uint64(time.Now().UnixNano()/100)
	return meow
}

// "A better solution is to obtain a 47-bit cryptographic quality random
// number and use it as the low 47 bits of the node ID, with the least
// significant bit of the first octet of the node ID set to one." - RFC 4122
func (u *uuidState) setNode() {
	if mac, err := defaultHWAddr(); err == nil {
		for b := 0; b < 6; b++ {
			u.Node[b] = mac[b]
		}
		return
	}

	var b [6]byte
	n, err := rand.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	if n == 0 {
		log.Fatal("No bytes read setting the node")
	}
	b[0] |= 0x01
	u.Node = b
}

func (u *uuidState) setClockSeq() error {
	var buf []byte = make([]byte, 2)
	_, err := rand.Read(buf[:])
	if err != nil {
		return err
	}
	u.ClockSeq = binary.BigEndian.Uint16(buf)
	return nil
}

// Returns hardware address if exists
func defaultHWAddr() (net.HardwareAddr, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []byte{}, err
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) >= 6 {
			return iface.HardwareAddr, nil
		}
	}
	return []byte{}, fmt.Errorf("No MAC address found")
}
