package enc

import (
	"goplay/config"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func TestNewJWT(t *testing.T) {
	key := config.MustLoad().GetJWTKey()
	token := jwt.New(jwt.SigningMethodHS256)
	s, err := token.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Signed string: %s\n", s)
}

func TestJWTWithClaims(t *testing.T) {
	id := uuid.New().String()
	key := config.MustLoad().GetJWTKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	s, err := token.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Signed string: %s\n", s)

	token, err = jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Token.Raw: %s\n", token.Raw)
	t.Logf("Token.Method: %v\n", token.Method)
	t.Logf("Token.Header: %v\n", token.Header)
	t.Logf("Token.Claims: %v\n", token.Claims)
	t.Logf("Token.Signature: %s\n", token.Signature)
	t.Logf("Token.Valid: %t\n", token.Valid)
}
