package httphandler

import (
	"net/http"
	"strconv"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/SashaMelva/smart_filter/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateUser(ctx *gin.Context) {
	var userCreaeter entity.UserCreater
	var user entity.User
	if err := ctx.ShouldBindJSON(&userCreaeter); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	accoountId := ctx.GetInt("accountId")

	years, err := pkg.GetAgeUser(user.DateBirthday)
	s.log.Debug(user)

	user = entity.User{
		AccountId:    accoountId,
		PhoneNumber:  userCreaeter.PhoneNumber,
		FirstName:    userCreaeter.FirstName,
		MiddelName:   userCreaeter.MiddelName,
		LastName:     userCreaeter.LastName,
		Age:          years,
		DateBirthday: userCreaeter.DateBirthday,
		AgeCategory:  pkg.GetAgeCategoryId(years),
	}

	id, err := s.app.CreateUser(&user)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) GetUser(ctx *gin.Context) {
	var user *entity.User
	var err error

	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	user, err = s.app.GetUserById(id)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	s.log.Debug(user)
	ctx.JSON(http.StatusOK, user)
}

func (s *Service) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = s.app.DeleteUser(id)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) UpdateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	s.log.Debug(user)

	err := s.app.UpdateUser(&user)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
