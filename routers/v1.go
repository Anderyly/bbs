package routers

import (
	v1 "bbs/app/controllers/server/v1"
	"github.com/gin-gonic/gin"
)

func V1Router(r *gin.Engine) {

	router := r.Group("/v1")
	router.POST("register", v1.Passport{}.Register)
	router.POST("login", v1.Passport{}.Login)

	router.POST("article/list", v1.Article{}.List)
}
