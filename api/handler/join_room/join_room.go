package join_room

import (
	"github.com/gin-gonic/gin"
	"github.com/losjalapenos/biletsele/api/router"
	"go.uber.org/fx"
)

type Result struct {
	fx.Out

	Handler router.Handler `group:"handlers"`
}

func joinRoom(_ *gin.Context) {
}

func Handler() Result {
	return Result{
		Handler: router.Handler{
			HTPPMethod:   "POST",
			RelativePath: "/join_room",
			Handler:      joinRoom,
		},
	}
}
