package chat

import "github.com/gin-gonic/gin"

type chatDeliveryImplementation struct {
	
}

func NewChatDelivery(router *gin.RouterGroup) *chatDeliveryImplementation {
	chatDelivery := &chatDeliveryImplementation{}

	router.GET("/", chatDelivery.Home)

	return chatDelivery
}

func (d *chatDeliveryImplementation) Home(c *gin.Context) {
	
}