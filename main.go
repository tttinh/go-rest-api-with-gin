package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/app/group"
	"github.com/tttinh/go-rest-api-with-gin/infra/config"
	"github.com/tttinh/go-rest-api-with-gin/infra/persistence"
	"github.com/tttinh/go-rest-api-with-gin/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Loading configuration.
	cfg := config.NewConfig()

	// Connecting DB.
	db := persistence.NewDB(cfg)

	// Setup Gin.
	gin.SetMode(cfg.Server.Mode)
	r := gin.Default()

	// Create application logic services.
	groupRepository := repository.NewGroupRepository(db)
	groupService := group.NewService(groupRepository)
	groupController := group.NewController(groupService)
	group.SetRoutes(r, groupController)

	// Start server.
	run(cfg, r)
}

func run(cfg config.Config, r *gin.Engine) {
	readTimeout := time.Duration(cfg.Server.ReadTimeout) * time.Second
	writeTimeout := time.Duration(cfg.Server.WriteTimeout) * time.Second
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           cfg.Server.Port,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	errs := make(chan error, 2)
	go func() {
		log.Printf("[info] start server on port %s", cfg.Server.Port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("terminated %v", <-errs)
}
