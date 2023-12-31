package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetSession(r *gin.Engine) *gin.Engine {

	store := cookie.NewStore([]byte("Andreyly"))
	r.Use(sessions.Sessions("session", store))

	return r

}
