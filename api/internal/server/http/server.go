package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/SashaMelva/smart_filter/internal/app"
	"github.com/SashaMelva/smart_filter/internal/config"
	"github.com/SashaMelva/smart_filter/internal/handler/httphandler"
	"github.com/SashaMelva/smart_filter/pkg"
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

	router.POST("/reg", handler.RegHendler)
	router.POST("/auth", handler.AuthHendler)

	protectedUser := router.Group("/user")
	protectedUser.Use(AuthMiddleware(log))
	{
		protectedUser.GET("/:id", handler.GetUser)
		protectedUser.POST("/", handler.CreateUser)
		protectedUser.PUT("/", handler.UpdateUser)
		protectedUser.DELETE("/:id", handler.DeleteUser)

		protectedUser.GET("/account/", handler.GetUserAccount)

	}

	protectedParent := router.Group("/children")
	protectedParent.Use(AuthMiddleware(log))
	{
		protectedParent.GET("/list/", handler.GetListChildren)
		protectedParent.POST("/link/:id", handler.AddGetChildren)

		protectedParent.PUT("/filters-gener/", handler.AddChildrenFilterGener)

		protectedParent.GET("/history-content-genre/:id/:date_start", handler.GetHistoryByCategoriesVideos)
	}

	videoFilters := router.Group("/filters")
	{
		videoFilters.GET("/chaild/:id", handler.GetChildrenFilter)
		videoFilters.GET("/age-category/:id", handler.GetFilterAgeCategory)
		// videoFilters.GET("/gener-category/:id", handler.GetGenre)
	}

	protectedVideo := router.Group("/video")
	{
		protectedVideo.POST("/chek", handler.ChekVideo)
		router.POST("/video", handler.AddNewVideo)
		router.PUT("/video", handler.UpdateVideo)
		router.PUT("/video-status", handler.UpdateStatusVideo)
		router.GET("/all-status", handler.GetAllStatus)
		router.GET("/all-age-category", handler.GetAllAgeCategory)
	}

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

func AuthMiddleware(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		log.Debug(token)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		sub, err := pkg.ParseAccessToken(token, "secretJWT")

		log.Debug(sub)
		if err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		intAcc, err := strconv.Atoi(sub)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Error(err)
			return
		}

		c.Set("accountId", intAcc)
		c.Next()
	}
}
