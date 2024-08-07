package main

import (
	routes "Task2/routers"
	"github.com/gin-gonic/gin"
	"encoding/gob"
)
func init() {
    gob.Register(gin.H{})
}

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*")

    routes.InitRouter(r)

    r.Run(":8080")
}
