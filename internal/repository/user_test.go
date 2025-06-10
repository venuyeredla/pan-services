package repository

import (
	"fmt"
	"testing"

	"github.com/venuyeredla/pan-services/internal/models"
)

func TestAuthenticate(t *testing.T) {
	IntializePool()
	defer ClosePool()
	euser, authError := Authenticate(models.AuthRequest{UserName: "venugopal@ecom.com", Password: "ecom#24"})
	if authError == nil {
		fmt.Println(euser)
	} else {
		t.Fail()
	}

}
