package main

import (
	"github.com/UjjwalMahar/Password-Gen/controllers"
	"github.com/UjjwalMahar/Password-Gen/initializer"
	"github.com/gin-gonic/gin"
)

func init(){
	initializer.LoadEnvVariables()
}

func main() {

	r := gin.Default()
	r.GET("/", controllers.GetRoot)
	r.POST("/generate-password", controllers.PostPasswordGen)

	r.Run()
	

}