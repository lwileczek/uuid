package main

import (
	"regexp"
	"testing"
)

//TOOD: Unit test GenerateUUID

//TestPseudoUUID Unit test for our Pseudo UUIDs
func TestPseudoUUID(t *testing.T) {
	exPseudoUUID, err := pseudoUUID()
	if err != nil {
		t.Error("Error Generating Pseudo UUID", err)
	}

	//Verify it is not nill as last check
	for _, b := range exPseudoUUID {
		if b != 0 {
			return
		}
	}
	t.Error("The UUID is the nil UUID")
}

func TestFormattingUUID(t *testing.T) {
	u, err := pseudoUUID()
	if err != nil {
		t.Error("Error Generating Pseudo UUID", err)
	}
	str := formatUUID(u)
	re, err := regexp.Compile("[A-Z0-9]{8}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{12}")
	if err != nil {
		t.Error("Error Generating Regular Expression", err)
	}
	if !re.MatchString(str) {
		t.Errorf("String produced did not match regular expression pattern:\nString:%s\n", str)
	}
}

func BenchmarkPseudoUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pseudoUUID()
	}
}
