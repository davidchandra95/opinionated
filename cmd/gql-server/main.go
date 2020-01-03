package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	host := "localhost"
	port := "7777"
	pathGQL := "/graphql"

	r := gin.Default()

	// setup a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	log.Println("Server running at http://" + host + ":" + port + pathGQL)
	log.Fatalln(r.Run(host + ":" + port))
}
