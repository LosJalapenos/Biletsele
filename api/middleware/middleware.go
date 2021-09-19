package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Provide(
	fx.Annotated{
		Name: "sessionMiddleware",
		Target: func() gin.HandlerFunc {
			store, _ := redis.NewStore(1, "tcp", "localhost:6379", "", []byte("secret"))
			return sessions.Sessions("token", store)
		},
	},
	fx.Annotated{
		Name: "authSessionMiddleware",
		Target: func() gin.HandlerFunc {
			return func(c *gin.Context) {
				if sessions.Default(c).Get("id") == nil {
					c.AbortWithStatus(http.StatusForbidden)
				}
			}
		},
	},
)
