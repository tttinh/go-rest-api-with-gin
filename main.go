package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/application/group"
	"github.com/tttinh/go-rest-api-with-gin/pkg/setting"
	"github.com/tttinh/go-rest-api-with-gin/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	setting.Setup()
	repositories := repository.New()
	groupService := group.NewService(repositories.Group)

	r := gin.Default()
	group.MakeHandler(r.Group("/api/v1/group"), groupService)

	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}

	errs := make(chan error, 2)
	go func() {
		log.Printf("[info] start http server listening 8080")
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("terminated %v", <-errs)
}
