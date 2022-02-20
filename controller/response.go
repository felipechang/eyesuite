package controller

import "github.com/gin-gonic/gin"

func Error(message string) *gin.H {
	return &gin.H{
		"data":  message,
		"error": true,
	}
}

func Success(data interface{}) *gin.H {
	return &gin.H{
		"data":  data,
		"error": false,
	}
}
