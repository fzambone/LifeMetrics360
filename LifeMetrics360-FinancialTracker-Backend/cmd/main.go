package main

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run()
}
