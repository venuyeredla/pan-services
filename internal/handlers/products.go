package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venuyeredla/pan-services/internal/models"
)

func GetProducts(gctx *gin.Context) {
	//val, ok := gctx.Params.Get("name")
	products := make([]models.Product, 0)
	products = append(products, models.Product{Name: "IPhone", Price: 700.5})
	products = append(products, models.Product{Name: "Pixel", Price: 600.5})
	products = append(products, models.Product{Name: "Samsung", Price: 500.5})
	gctx.JSON(http.StatusOK, products)

}
