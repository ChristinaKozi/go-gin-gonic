package main

import (
	"github.com/ChristinaKozi/go-gin-gonic/controllers"
	"github.com/ChristinaKozi/go-gin-gonic/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controllers.ConnectToDB()

	routes.Routes(router)

	router.Run(":8080")
}
