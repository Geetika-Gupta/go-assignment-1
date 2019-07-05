package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	shipperValidator "project/validators"
	routes "project/routes"
)

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		shipperValidator.RegisterShipperValidations(v)
	}
	routes.InitializeRoutes(router)
	router.Run()
}