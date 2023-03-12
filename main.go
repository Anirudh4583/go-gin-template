package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Anirudh4583/go-gin-template/models"
	"github.com/Anirudh4583/go-gin-template/pkg/setting"
	util "github.com/Anirudh4583/go-gin-template/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	util.Setup()
	models.Setup()
}

func main() {
	println("Hello GO API!")

	gin.SetMode(setting.Config.ServerRunMode)

	port := fmt.Sprintf(":%d", setting.Config.ServerHttpPort)
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "health check passed!")
	})

	var err error

	readTimout, err := time.ParseDuration(setting.Config.ServerReadTimeout + "s")
	if err != nil {
		log.Fatalf("cannot parse server read timeout: %v", err)
	}
	writeTimeout, err := time.ParseDuration(setting.Config.ServerWriteTimeout + "s")
	if err != nil {
		log.Fatalf("cannot parse server write timeout: %v", err)
	}

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
