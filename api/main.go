package main

import (
	"github.com/gin-gonic/gin"
	"github.com/losjalapenos/biletsele/api/handler"
	"github.com/losjalapenos/biletsele/api/middleware"
	"github.com/losjalapenos/biletsele/api/router"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		middleware.Module,
		handler.Module,
		router.Module,
		fx.Invoke(func(r *gin.Engine) {
			_ = r.Run(":8080")
		}),
	).Run()
}
