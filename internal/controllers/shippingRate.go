package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/candles_api/internal/services"
)

type ShippingRateController struct {
	shippingRateService services.ShippingRateService
}

func NewShippingRateController(shippingRateService services.ShippingRateService) *ShippingRateController {
	return &ShippingRateController{shippingRateService: shippingRateService}
}

func (sc *ShippingRateController) GetShippingRate(c *gin.Context) {
	zipCode := c.Param("zipCode")
	quantityStr := c.Query("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	rate, err := sc.shippingRateService.GetShippingRate(zipCode, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"zipCode": zipCode, "quantity": quantity, "rate": rate})
}
