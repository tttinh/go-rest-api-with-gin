package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tttinh/go-rest-api-with-gin/app/group"
	"github.com/tttinh/go-rest-api-with-gin/infra/config"
	"github.com/tttinh/go-rest-api-with-gin/infra/logger"
	"github.com/tttinh/go-rest-api-with-gin/infra/persistence"
	"github.com/tttinh/go-rest-api-with-gin/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Loading configuration.
	cfg := config.NewConfig()

	// Init logger
	logger.Initialize(cfg.Server.Mode)

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
		logger.Info("starting server on port ", cfg.Server.Port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Info("server stopped: ", <-errs)
}
