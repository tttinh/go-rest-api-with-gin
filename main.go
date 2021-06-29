package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/application/group"
	"github.com/tttinh/go-rest-api-with-gin/repository"
	"github.com/tttinh/go-rest-api-with-gin/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	appSetting := setting.Load("development")
	repositories := repository.New(appSetting.Database)
	groupService := group.NewService(repositories.Group)
	groupController := group.NewController(groupService)

	gin.SetMode(appSetting.Server.Mode)
	r := gin.Default()

	groupController.SetRoutes(r.Group("/api/v1/group"))

	run(appSetting, r)
}

func run(appConfig setting.Setting, r *gin.Engine) {
	readTimeout := time.Duration(appConfig.Server.ReadTimeout) * time.Second
	writeTimeout := time.Duration(appConfig.Server.WriteTimeout) * time.Second
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           appConfig.Server.Port,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	errs := make(chan error, 2)
	go func() {
		log.Printf("[info] start server on port %s", appConfig.Server.Port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("terminated %v", <-errs)
}
