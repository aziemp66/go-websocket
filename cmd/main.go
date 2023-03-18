package main

import (
	chatDlv "github.com/aziemp66/go-websocket/internal/delivery/chat"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	goViewConfig := goview.Config{
		Root:         "web/views",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: false,
		Delims: goview.Delims{
			Left:  "{{",
			Right: "}}",
		},
	}

	viewEngine := ginview.New(goViewConfig)

	router.HTMLRender = viewEngine

	router.Static("/assets", "./web/static/assets")

	chatGroup := router.Group("/")
	chatDlv.NewChatDelivery(chatGroup)

	router.Run(":8080")
}
