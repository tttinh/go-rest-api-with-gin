package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/application/group"
	"github.com/tttinh/go-rest-api-with-gin/config"
	"github.com/tttinh/go-rest-api-with-gin/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	appConfig := config.Load("development")
	repositories := repository.New(appConfig.Database)
	groupService := group.NewService(repositories.Group)
	groupController := group.NewController(groupService)

	gin.SetMode(appConfig.Server.Mode)
	r := gin.Default()

	groupController.SetRoutes(r.Group("/api/v1/group"))

	run(appConfig, r)
}

func run(appConfig config.Config, r *gin.Engine) {
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
		log.Printf("[info] start httperror server listening 8080")
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("terminated %v", <-errs)
}
