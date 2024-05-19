package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
	"github.com/sergey4qb/rate/internal/core/ports"
	"github.com/sergey4qb/rate/internal/core/service/rates_subscribers_service"
	"net/http"
)

type RatesSubscribersController struct {
	service ports.RatesSubscribersService
}

func NewRatesSubscribersController(service ports.RatesSubscribersService) *RatesSubscribersController {
	return &RatesSubscribersController{service: service}
}

func (controller *RatesSubscribersController) SaveSubscriberMail(c *gin.Context) {
	var req subscriber.Subscriber
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	err := controller.service.SaveSubscriber(c, req)
	if errors.Is(err, rates_subscribers_service.ErrSubscriberAlreadyExist) {
		respondWithError(c, http.StatusConflict, err)
		return
	}
	if errors.Is(err, subscriber.ErrInvalidEmail) {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
