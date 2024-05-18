package httphandler

import (
	"net/http"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) AddNewVideo(ctx *gin.Context) {
	var video entity.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err := s.app.AddNewVideo(video)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}

func (s *Service) ChekVideo(ctx *gin.Context) {
	var video entity.VideoCheker

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err := s.app.ChekVideo(video)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}

func (s *Service) GetAllStatus(ctx *gin.Context) {
	var user *entity.StatusVideos
	var err error

	user, err = s.app.GetAllStatusVideo()

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	s.log.Debug(user)
	ctx.JSON(http.StatusOK, user)
}

func (s *Service) GetAllAgeCategory(ctx *gin.Context) {
	var user *entity.AgeCategores
	var err error

	user, err = s.app.GetAllAgeCategoryVideo()

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	s.log.Debug(user)
	ctx.JSON(http.StatusOK, user)
}
