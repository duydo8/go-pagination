package main

import (
	"github.com/gin-gonic/gin"
	"pagination/controllers"
	"pagination/initializers"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	initializers.GenerateModel()
}
func main() {
	r := gin.Default()
	r.GET("findByPageable", controllers.FindByPageable)
	r.Run()

}
