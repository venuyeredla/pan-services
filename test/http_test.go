package wtest

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/venuyeredla/pan-services/internal/models"
)

const (
	DOMAIN_NAME = "http://localhost:2024"
)

func TestAuthEndpoint(t *testing.T) {
	/*
			{
		    "username":"",
		    "password" : ""
		    }
	*/
	requesBody := &models.AuthRequest{UserName: "venugopal@ecom.com", Password: "ecom#24"}
	bytearr, error := json.Marshal(requesBody)

	if error != nil {
		log.Default().Println("Error in marshalling strcut", error.Error())
	}
	req, _ := http.NewRequest("POST", DOMAIN_NAME+"/api/v1/auth/signin", bytes.NewBuffer(bytearr))
	response, h_error := http.DefaultClient.Do(req)
	if h_error == nil {
		bodyBytes, _ := io.ReadAll(response.Body)
		var resp models.AuthResponse
		json.Unmarshal(bodyBytes, &resp)
		log.Default().Println(resp)

	} else {
		log.Default().Println("Error in connecting api")
	}
	defer response.Body.Close()

}
