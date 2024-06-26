package httphandler

import (
	"net/http"
	"strconv"

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

	ok, err := s.app.ChekVideo(video)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	if ok {
		ctx.JSON(http.StatusOK, gin.H{"filters": "OK"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"filters": "Error"})
	}
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

func (s *Service) UpdateVideo(ctx *gin.Context) {
	var video entity.VideoId

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err := s.app.UpdateVideo(video)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}

func (s *Service) UpdateStatusVideo(ctx *gin.Context) {
	var video entity.VideoIdStatus

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err := s.app.UpdateStatusVideo(video)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}

func (s *Service) GetFilterAgeCategory(ctx *gin.Context) {
	var filters *entity.FilterAgeCategores
	var err error

	filters, err = s.app.GetFiltersAgeCategoryVideo()

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	s.log.Debug(filters)
	ctx.JSON(http.StatusOK, filters)
}

func (s *Service) GetHistoryByCategoriesVideos(ctx *gin.Context) {
	var procent *entity.ProcentByCategoresUser
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	date_start := ctx.Params.ByName("date_start")

	procent, err = s.app.GetHistoryByCategoriesVideos(id, date_start)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procent)
}

func (s *Service) GetResoursesFilter(ctx *gin.Context) {
	var procent *entity.ProcentByCategoresUser
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	procent, err = s.app.GetHistoryByResourcesVideos(id)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procent)
}
