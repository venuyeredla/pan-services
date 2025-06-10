package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venuyeredla/pan-services/internal/models"
	"github.com/venuyeredla/pan-services/internal/repository"
)

type CustomerAPI struct {
}

func (capi *CustomerAPI) CustInfo(c *gin.Context) {
	customer, authError := repository.Authenticate(models.AuthRequest{})
	if authError != nil {
		// json.NewEncoder(w).Encode(customer)
		c.JSON(http.StatusOK, customer)
		//fmt.Fprintf(w, "Hello, world = %s", "Coder")
	}

}
