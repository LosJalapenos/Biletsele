package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	HTPPMethod   string
	RelativePath string
	Handler      gin.HandlerFunc
}

type params struct {
	fx.In

	SessionMiddleware gin.HandlerFunc `name:"sessionMiddleware"`
	ConnectHandler    Handler         `name:"connectHandler"`

	AuthSessionMiddleware gin.HandlerFunc `name:"authSessionMiddleware"`
	Handlers              []Handler       `group:"handlers"`
}

var Module = fx.Provide(
	func(p params) *gin.Engine {
		r := gin.Default()

		r.Use(p.SessionMiddleware)
		r.Handle(p.ConnectHandler.HTPPMethod, p.ConnectHandler.RelativePath, p.ConnectHandler.Handler)

		r.Use(p.AuthSessionMiddleware)
		for _, h := range p.Handlers {
			r.Handle(h.HTPPMethod, h.RelativePath, h.Handler)
		}

		return r
	},
)
