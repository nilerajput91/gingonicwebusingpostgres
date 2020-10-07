package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nilerajput91/gingonic/configs"
	routes "github.com/nilerajput91/gingonic/routes"
)

func main() {

	configs.Connect()

	//Init Router
	router := gin.Default()

	//Router and handlers/Endpoints

	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
