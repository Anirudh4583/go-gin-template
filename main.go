package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anirudh4583/go-gin-template/pkg/setting"
	util "github.com/Anirudh4583/go-gin-template/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	util.Setup()
}

func main() {
	println("Hello GO API!")

	gin.SetMode(setting.Config.ServerSetting.RunMode)

	port := fmt.Sprintf(":%d", setting.Config.ServerSetting.HttpPort)
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "health check passed!")
	})
	readTimout := setting.Config.ServerSetting.ReadTimeout
	writeTimeout := setting.Config.ServerSetting.WriteTimeout

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", port)
	s.ListenAndServe()
}
