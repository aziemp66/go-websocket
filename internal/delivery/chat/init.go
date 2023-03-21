package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatDeliveryImplementation struct {
}

func NewChatDelivery(router *gin.RouterGroup) *chatDeliveryImplementation {
	chatDelivery := &chatDeliveryImplementation{}

	router.GET("/", chatDelivery.Home)
	router.GET("/page", chatDelivery.Page)
	router.GET("/chat", chatDelivery.Chat)

	return chatDelivery
}

func (d *chatDeliveryImplementation) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Home",
		"add": func(a, b int) int {
			return a + b
		},
	})
}

func (d *chatDeliveryImplementation) Page(c *gin.Context) {
	c.HTML(http.StatusOK, "page.html", gin.H{
		"title": "Page",
	})
}

func (d *chatDeliveryImplementation) Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat", gin.H{
		"title": "Chat",
	})
}
