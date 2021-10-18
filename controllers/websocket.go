package controllers

import (
	"net/http"
	"gopkg.in/olahol/melody.v1"
	"github.com/gin-gonic/gin"
	"app/utils"

)

func WebSocket(c *gin.Context){

	if !utils.CheckToken(c.Request.Header["Token"][0]){
	
		c.String(400,"Bunday token yo'q")
		return
	}else {
		webSocketRouter := melody.New()

		webSocketRouter.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }


		webSocketRouter.HandleRequest(c.Writer, c.Request)

		webSocketRouter.HandleMessage(func (s *melody.Session, m []byte) {
			webSocketRouter.Broadcast(m)
		})
	}
}