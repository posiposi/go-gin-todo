package controller

import (
	"github.com/gin-gonic/gin"
)

const successResponseCode = 200

func HelloWorld(c *gin.Context) {
	responseMessage := "Hello World!"
	c.String(successResponseCode, responseMessage)
}

func PrintTestWord(c *gin.Context) {
	responseMessage := "これはテスト用出力です。"
	c.String(successResponseCode, responseMessage)
}
