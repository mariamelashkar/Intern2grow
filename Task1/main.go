package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"task_1/routes"


)

func main() {
	router := gin.Default()
	routers.InitRoutes(router)
	router.Run("localhost:8000")
	log.Println("Server started at :8000")
	
}
