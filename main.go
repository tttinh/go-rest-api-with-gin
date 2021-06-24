package main

import (
	"example.com/demo/application/group"
	"example.com/demo/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	r := gin.Default()

	groupRepository := repository.NewGroupRepository(nil)
	groupService := group.NewService(groupRepository)
	groupEndpoints := group.NewEndpoints(groupService)
	groupRouters := r.Group("/api/v1/group")
	group.SetupHandler(groupRouters, groupEndpoints)

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
