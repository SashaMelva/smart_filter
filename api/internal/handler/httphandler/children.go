package httphandler

import (
	"net/http"
	"strconv"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetListChildren(ctx *gin.Context) {
	var list *entity.ChilgrenLists
	accoountId := ctx.GetInt("accountId")

	s.log.Debug(accoountId)
	list, err := s.app.GetListChildren(accoountId)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"list": list})
}

func (s *Service) AddGetChildren(ctx *gin.Context) {

	accoountId := ctx.GetInt("accountId")
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = s.app.AddListChildren(accoountId, id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}

func (s *Service) GetChildrenFilter(ctx *gin.Context) {
	var filters *entity.Fileters
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	filters, err = s.app.GetFiltersByChaild(id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"filters": filters.Filters})
}

func (s *Service) AddChildrenFilterGener(ctx *gin.Context) {
	var filters entity.FilterGeners

	if err := ctx.ShouldBindJSON(&filters); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err := s.app.AddFiltersGanreForChaild(&filters)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ``)
}
func (s *Service) DeleteChildrenFilterGener(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, ``)
}
