package httphandler

import (
	"net/http"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) AuthHendler(ctx *gin.Context) {
	var account entity.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.app.Auth(&account)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) RegHendler(ctx *gin.Context) {
	var account entity.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	s.log.Debug(account)

	id, err := s.app.CreateAccount(&account)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}
