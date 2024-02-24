package router

import (
	"net/http"

	"DockerGin/config"
	"DockerGin/src/infrastructure/db"

	"github.com/gin-gonic/gin"
)

func Dispatch(sqlhandler *db.SQLHandler) {
	conf := config.Get()

	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.Run(":" + conf.Port)
}
