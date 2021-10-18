package main

import(
	c "app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"

)

func Home(ctx *gin.Context) {
	ctx.JSON(200,gin.H {
		"message": "PONG",
	})
}

func main() {

	r := gin.Default()

	r.GET("/", Home)
	r.POST("/users",c.Login)
	r.POST("/register",c.Register)
	r.POST("/change",c.Change)
	r.GET("/echo", c.WebSocket)
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.Run()
}