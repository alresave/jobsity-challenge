package handler

import (
	"github.com/gin-gonic/gin"
	"jobsity-challenge/chat-service/internal/chat"
	"jobsity-challenge/common/middleware"
)

func SetupRouter(svr chat.Service) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	router.Static("/front", "../frontend/chat/")
	/*router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "../frontend/chat/index.html")
	})*/
	router.Group("/chat")
	{
		v1Group := router.Group("/chat/v1")
		{
			v1Group.GET("/messages", svr.GetMessages)
			v1Group.GET("/rooms", svr.GetRooms)
			v1Group.GET("/room_users", svr.GetRoomUsers)
			v1Group.GET("/ticket", svr.GetChatTicket)
			v1Group.GET("/ws", svr.HandleChat)
		}
	}
	return router
}
