package http

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/SashaMelva/smart_filter/internal/app"
	"github.com/SashaMelva/smart_filter/internal/config"
	"github.com/SashaMelva/smart_filter/internal/handler/httphandler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
	log *zap.SugaredLogger
}

func NewServer(log *zap.SugaredLogger, app *app.App, config *config.ConfigHttpServer) *Server {
	log.Debug("URL: " + config.Host + ":" + config.Port)
	router := gin.Default()
	handler := httphandler.NewHendler(log, app)

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("Hellow world)")
	})

	router.GET("/user/:id", handler.GetUser)
	router.POST("/user/", handler.CreateUser)
	router.PUT("/user/", handler.UpdateUser)
	router.DELETE("/user/:id", handler.DeleteUser)

	// router.GET("/account-chaild/", handler.GetAccountsChailds)
	// router.POST("/account-chaild/", handler.linkingСhildsAccount)

	return &Server{
		srv: &http.Server{
			Addr:    config.Host + ":" + config.Port,
			Handler: router,
		},
		log: log,
	}
}

func (s *Server) Start(ctx context.Context) {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Fatalf("listen: %s\n", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(1)
}
