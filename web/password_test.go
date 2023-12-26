package web

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPassword(t *testing.T) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("12345678"), 14)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", bytes)

	err = bcrypt.CompareHashAndPassword(bytes, []byte("12345678"))
	if err != nil {
		t.Log("Password not matched")
		t.Fatal(err)
	} else {
		t.Log("Password matched")
	}

}
