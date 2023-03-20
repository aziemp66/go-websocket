package chat

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	wsCommon "github.com/aziemp66/go-websocket/common/websocket"
)

type chatDeliveryImplementation struct {
}

func NewChatDelivery(router *gin.RouterGroup) *chatDeliveryImplementation {
	chatDelivery := &chatDeliveryImplementation{}

	router.GET("/", chatDelivery.Home)
	router.GET("/page", chatDelivery.Page)
	router.GET("/websocket", chatDelivery.Websocket)

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

func (d *chatDeliveryImplementation) Websocket(c *gin.Context) {
	ws, err := wsCommon.Upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Client connected to endpoint: /chat")

	var message wsCommon.WsJsonResponse

	message.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "chat", gin.H{
		"title": "Chat",
	})
}
