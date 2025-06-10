package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/venuyeredla/pan-services/internal/models"
	"github.com/venuyeredla/pan-services/internal/repository"
)

const (
	ECRYPT_KEY  = "conceal#24"
	ECRYPT_ALGO = "HS256"
)

func Authenticate(gctx *gin.Context) {
	//val, ok := gctx.Params.Get("name")
	var authReq models.AuthRequest
	json.NewDecoder(gctx.Request.Body).Decode(&authReq)
	log.Default().Printf("Login request : %s", &authReq)
	authUser, authError := repository.Authenticate(authReq)
	if authError == nil {
		claims := jwt.MapClaims{
			"first_name": authUser.Firstname,
			"last_name":  authUser.Lastname,
			"roles":      authUser.Roles,
		}
		signedToken := GenJwtToken(authUser.Email, claims)
		authResp := &models.AuthResponse{Token: signedToken, Algo: ECRYPT_ALGO}
		gctx.JSON(http.StatusOK, authResp)
	} else {
		eresp := &models.ErrorResponse{Msg: "Invalid credentials"}
		gctx.JSON(http.StatusUnauthorized, eresp)
	}
}

func SignUp(c *gin.Context) {
	var customer models.User
	json.NewDecoder(c.Request.Body).Decode(&customer)
	c.String(http.StatusOK, "Success fully signed up", customer.Firstname, customer.Lastname)
}

// JWT methods

func GenJwtToken(subject string, claims jwt.MapClaims) string {
	timeNow := time.Now()
	claims["sub"] = subject
	claims["nbf"] = jwt.NewNumericDate(timeNow)
	claims["iat"] = jwt.NewNumericDate(timeNow)
	claims["exp"] = jwt.NewNumericDate(timeNow.AddDate(0, 1, 0))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(ECRYPT_KEY))
	return signedToken
}

func ParseToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(ECRYPT_KEY), nil
	})
	if err != nil {
		log.Default().Println("error in token parsing")
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	} else {
		return nil
	}
}

func IsValidToken(claims jwt.MapClaims) bool {
	numericDate, error := claims.GetExpirationTime()
	if error == nil {
		if time.Now().Before(numericDate.Time) {
			return true
		} else {
			return false
		}
	}
	return false
}
