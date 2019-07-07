package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Geetika-Gupta/go-assignment-1/custom_validators"
	"github.com/Geetika-Gupta/go-assignment-1/routes"
)

func main() {
	router := gin.Default()
	custom_validators.RegisterValidations()
	routes.InitializeRoutes(router)
	router.Run()
}