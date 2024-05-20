package api

import (
	"go-gin-todo/internal/interface/http/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// v1APIを定義
	v1 := router.Group("/api/v1")
	{
		// テスト用出力
		v1.GET("/hello", controller.HelloWorld)

		// todo一覧表示
		v1.GET("/todos", controller.PrintTestWord)
	}

	return router
}
