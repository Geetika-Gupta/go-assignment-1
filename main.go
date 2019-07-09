package main

import (
	cv "github.com/Geetika-Gupta/go-assignment-1/custom_validator"
	"github.com/Geetika-Gupta/go-assignment-1/shipper"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cv.RegisterValidations()
	shipper.InitializeRoutes(router)
	router.Run()
}
