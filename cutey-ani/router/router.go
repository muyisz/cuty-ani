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
	router.GET("/dash", handler.GetDash)
	router.GET("/info", handler.GetInfo(db))
	return router
}
