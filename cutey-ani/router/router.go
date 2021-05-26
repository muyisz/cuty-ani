package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/cutey-ani/data"
	"github.com/muyisz/cutey-ani/handler"
)

func InitRouter(db *data.MySQL) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/html/*")
	router.StaticFS("/views", http.Dir("./views"))
	router.GET("/", handler.GetHome)
	router.POST("/login", handler.PostLoginData(db))
	router.POST("/register", handler.PostRegisterData(db))
	router.GET("/login", handler.GetLogin)
	router.GET("/info", handler.GetInfo(db))
	router.POST("/up_photo", handler.PostPhoto(db))
	router.GET("/getphoto", handler.GetPhoto(db))
	{
		router.POST("/postRoom", handler.CheckCookie(db), handler.PostChatRoom(db))
		router.GET("/dash", handler.CheckCookie(db), handler.GetDash)
		router.GET("/getRoom", handler.CheckCookie(db), handler.GetChatRoom(db))
		router.GET("/message", handler.CheckCookie(db), handler.GetRoom)
	}
	return router
}
