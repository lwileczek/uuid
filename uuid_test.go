package uuid

import (
	"regexp"
	"testing"
)

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
	str := FormatUUID(u)
	re, err := regexp.Compile("[A-Z0-9]{8}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{12}")
	if err != nil {
		t.Error("Error Generating Regular Expression", err)
	}
	if !re.MatchString(str) {
		t.Errorf("String produced did not match regular expression pattern:\nString:%s\n", str)
	}
}

func TestGenerateUUID(t *testing.T) {
	u, err := GenerateUUID("v1", false)
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

	u, err = GenerateUUID("v4", false)
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
	u, err = GenerateUUID("pseudo", false)
	if err != nil {
		t.Error("Error Generating Pseudo UUID", err)
	}
	if u == Nil {
		t.Error("Error, pseudo generator returned null UUID")
	}
	u, err = GenerateUUID("null", false)
	if err != nil {
		t.Error("Error Generating Pseudo UUID", err)
	}
	if u != Nil {
		t.Error("NULL UUID generator did not create null UUID:", u)
	}
	u, err = GenerateUUID("GiveMeError", false)
	if err == nil {
		t.Error("Generate UUID is not tossing error on random input")
	}
}

func BenchmarkPseudoUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pseudoUUID()
	}
}
