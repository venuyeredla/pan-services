package handlers

import (
	"testing"

	jwt "github.com/golang-jwt/jwt/v5"
)

func TestJwtToken(t *testing.T) {

	claims := jwt.MapClaims{
		"first_name": "venu",
		"last_name":  "gopal",
	}
	token := GenJwtToken("venugopal@ecom.com", claims)
	parsedClaims := ParseToken(token)
	isValid := IsValidToken(parsedClaims)
	if !isValid {
		t.Errorf("Invalid token")
		t.FailNow()
	}
}
