package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/rate/internal/core/ports"
	"net/http"
)

type ExchangeRateController struct {
	service ports.ExchangeRatesService
}

func NewExchangeRateController(service ports.ExchangeRatesService) *ExchangeRateController {
	return &ExchangeRateController{service: service}
}

func (controller *ExchangeRateController) GetUAHUSDExchangeRate(c *gin.Context) {
	u, err := controller.service.GetUAHUSDExchangeRate(c)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	response := UAHUSDRateResponse{Rate: u}
	c.JSON(http.StatusOK, response)
}
