package server

import (
	"github.com/davidchandra95/opinionated/pkg/utils"
	"log"

	"github.com/davidchandra95/opinionated/internal/handlers"
	"github.com/gin-gonic/gin"
)

var HOST, PORT string

func init() {
	HOST = utils.MustGet("GQL_SERVER_HOST")
	PORT = utils.MustGet("GQL_SERVER_PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running at http://" + HOST + ":" + PORT )
	log.Fatalln(r.Run(HOST + ":" + PORT))
}