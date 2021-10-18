package main

import(
	c "app/controllers"
	"github.com/gin-gonic/gin"

)

func main() {

	r := gin.Default()

	r.POST("/login",c.Login)
	r.POST("/register",c.Register)
	r.POST("/change",c.Change)
	r.GET("/echo", c.WebSocket)

	r.Run()
}