package main

import (
	"log"
	"net/http"

	chatDlv "github.com/aziemp66/go-websocket/internal/delivery/chat"

	wsCommon "github.com/aziemp66/go-websocket/common/websocket"

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

	// serving websocket
	router.GET("/ws/:id", wsCommon.WsEndpoint)

	router.StaticFS("/public", http.Dir("web/public"))

	chatGroup := router.Group("/")
	chatDlv.NewChatDelivery(chatGroup)

	go wsCommon.ListenToWsChannel()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router listen error: %s\n", err)
	}
}
