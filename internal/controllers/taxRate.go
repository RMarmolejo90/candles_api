package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/candles_api/internal/services"
)

type TaxRateController struct {
	taxRateService services.TaxRateService
}

func NewTaxRateController(taxRateService services.TaxRateService) *TaxRateController {
	return &TaxRateController{taxRateService: taxRateService}
}

func (tc *TaxRateController) GetTaxRate(c *gin.Context) {
	zipCode := c.Param("zipCode")
	rate, err := tc.taxRateService.GetTaxRate(zipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"zipCode": zipCode, "rate": rate})
}
