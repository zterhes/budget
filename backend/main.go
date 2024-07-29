package main

import (
	"zterhes/budget/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/transactions/upload", controllers.UploadTransactions)
	router.Run("localhost:8080")
}
