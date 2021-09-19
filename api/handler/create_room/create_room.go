package create_room

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/losjalapenos/biletsele/api/router"
	"go.uber.org/fx"
	"net/http"
)

type Result struct {
	fx.Out

	Handler router.Handler `group:"handlers"`
}

func createRoom(c *gin.Context) {
	session := sessions.Default(c)

	c.JSON(http.StatusOK, gin.H{
		"id":   session.Get("id"),
		"name": session.Get("name"),
	})
}

func Handler() Result {
	return Result{
		Handler: router.Handler{
			HTPPMethod:   "GET",
			RelativePath: "/create_room",
			Handler:      createRoom,
		},
	}
}
