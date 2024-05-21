package httphandler

import (
	"net/http"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Auth
// @Tags auth
// @Description create account
// @ID auth-account
// @Accept  json
// @Produce  json
// @Param input body entity.Account true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /auth [post]

func (s *Service) AuthHendler(ctx *gin.Context) {
	var account entity.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	s.log.Debug(account)

	id, err := s.app.Auth(&account)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

// @Summary Registration
// @Tags reg
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body entity.Account true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /reg [post]

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
