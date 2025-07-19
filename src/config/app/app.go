package config_app

import (
	"context"
	"fmt"
	config_db "go-template/src/config/db"
	config_general "go-template/src/config/general"
	"go-template/src/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func NewServer(
	server *gin.Engine,
	conf config_general.AllConfig,
	db config_db.Connection,
	logger *logrus.Logger,
	validator *validator.Validate,
) *http.Server {
	// init route
	router.InitRoute(server, conf, db, logger, validator)
	port := conf.HTTPConfig.HttpPort
	fmt.Printf("RUNNING IN PORT %s\n", port)
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: server.Handler(),
	}
}

func StartService(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		fmt.Println("SERVICE START")
	}()

	gracefullShutdown(srv)

}

func gracefullShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
