package connect

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/losjalapenos/biletsele/api/router"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Result struct {
	fx.Out

	Handler router.Handler `name:"connectHandler"`
}

type connectRequest struct {
	Name string `json:"name" binding:"required,alpha,min=1,max=16"`
}

func connect(c *gin.Context) {
	session := sessions.Default(c)

	var connectRequest connectRequest
	if err := c.ShouldBindJSON(&connectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
		})
	}

	session.Set("id", fmt.Sprintf("%x", md5.Sum([]byte(connectRequest.Name+time.Now().String()))))
	session.Set("name", connectRequest.Name)

	if session.Save() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
		})
	}
}

func Handler() Result {
	return Result{
		Handler: router.Handler{
			HTPPMethod:   "POST",
			RelativePath: "/connect",
			Handler:      connect,
		},
	}
}
