package main

import (
	"zterhes/budget/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/uploadTransactions", services.UploadTransactions)
	router.Run("localhost:8080")
}
