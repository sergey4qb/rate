package router

import (
	"github.com/sergey4qb/rate/internal/core/delivery/http/controllers"
	"github.com/sergey4qb/rate/internal/core/ports"
	"time"
)
import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

type Router struct {
	gin                        *gin.Engine
	ExchangeRateController     *controllers.ExchangeRateController
	RatesSubscribersController *controllers.RatesSubscribersController
}

func NewRouter(
	gin *gin.Engine,
	exchangeRates ports.ExchangeRatesService,
	ratesSubscriber ports.RatesSubscribersService,
) *Router {
	r := &Router{
		gin: gin,
	}
	r.ExchangeRateController = controllers.NewExchangeRateController(exchangeRates)
	r.RatesSubscribersController = controllers.NewRatesSubscribersController(ratesSubscriber)
	r.setupCORS()
	r.setupRoutes()
	return r
}

func (r *Router) Run(port string) error {
	return r.gin.Run(port)
}
func (r *Router) setupCORS() {
	r.gin.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://localhost:8080", "http://localhost:8080", "https://1c08-31-128-77-167.ngrok-free.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "https://95a0-31-128-77-167.ngrok-free.app" || origin == "https://localhost:8080" || origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func (r *Router) setupRoutes() {
	r.setupExchangeRateRoutes()
	r.setupSubscribersRoutes()
}

func (r *Router) setupExchangeRateRoutes() {
	// group := r.gin.Group("/exchange_rate")
	// group.GET("/rate", r.ExchangeRateController.GetUAHUSDExchangeRate)
	r.gin.GET("/rate", r.ExchangeRateController.GetUAHUSDExchangeRate)
}

func (r *Router) setupSubscribersRoutes() {
	// group := r.gin.Group("")
	r.gin.POST("/subscribe", r.RatesSubscribersController.SaveSubscriberMail)
}
