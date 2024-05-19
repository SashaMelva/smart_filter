package app

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (a *App) ChekVideo(video entity.VideoCheker) (bool, error) {
	ok, err := a.storage.ChekVideo(video.UrlVideo)
	a.log.Debug("DB ", ok)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	if ok {
		a.log.Debug("Достаем видео и пользователя")
		a.log.Debug(video)
		user, err := a.storage.GetUserByIdAccount(video.IdAccount)
		a.log.Debug(user)
		if err != nil {
			return false, err
		}
		videoNew, err := a.storage.GetVideoByUrl(video.UrlVideo)
		a.log.Debug(videoNew)
		if err != nil {
			return false, err
		}

		a.log.Debug(user.AgeCategory < videoNew.AgeCategoryId, user.AgeCategory, videoNew.AgeCategoryId)
		if user.AgeCategory < videoNew.AgeCategoryId || videoNew.AgeCategoryId == 0 {
			return false, nil
		}

		arrGaner := strings.Split(strings.TrimSpace(user.GenersIds), ",")
		a.log.Debug(arrGaner, len(arrGaner))

		a.log.Debug(chekTrue(user.GenersIds, videoNew.GenerId), videoNew.GenerId, user.GenersIds)
		if chekTrue(user.GenersIds, videoNew.GenerId) {
			datetime := time.Now()
			err := a.storage.AddHistory(videoNew.Id, user.AccountId, datetime)

			if err != nil {
				return false, err
			}

			return true, nil
		}

		return false, nil
	}

	err = a.storage.AddNewVideo(video.UrlVideo, "")

	if err != nil {
		a.log.Error(err)
	}

	return false, nil
}

func chekTrue(ids string, videId int) bool {
	arrGaner := strings.Split(strings.TrimSpace(ids), ",")
	if len(arrGaner) == 1 {
		intW, err := strconv.Atoi(arrGaner[0])
		if err != nil {
			return true
		}
		if intW == videId {
			return true
		} else {
			return false
		}
	}
	for i := range arrGaner {
		intW, err := strconv.Atoi(arrGaner[i])

		if err != nil {
			return false
		}

		if intW == videId {
			return true
		}
	}

	return false
}

func (a *App) AddNewVideo(video entity.Video) error {

	ok, err := a.storage.ChekVideo(video.Url)

	if err != nil {
		a.log.Error(err)
		return err
	}

	if ok {
		return errors.New("Уже существует видео с этим url")
	}

	err = a.storage.AddNewVideo(video.Url, video.Name)

	if err != nil {
		a.log.Error(err)
	}

	return err
}

func (a *App) GetAllStatusVideo() (*entity.StatusVideos, error) {
	var list *entity.StatusVideos
	list, err := a.storage.GetStatusVideos()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}

func (a *App) GetAllAgeCategoryVideo() (*entity.AgeCategores, error) {
	var list *entity.AgeCategores
	list, err := a.storage.GetAllAgeCategoryVideos()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}

func (a *App) UpdateVideo(video entity.VideoId) error {
	err := a.storage.UpdateVideo(video)

	if err != nil {
		a.log.Error(err)
	}

	return err
}
func (a *App) UpdateStatusVideo(video entity.VideoIdStatus) error {
	err := a.storage.UpdateStatusVideo(video)

	if err != nil {
		a.log.Error(err)
	}

	return err
}

func (a *App) GetFiltersAgeCategoryVideo() (*entity.FilterAgeCategores, error) {
	var list *entity.FilterAgeCategores
	list, err := a.storage.GetFilterAgeCategory()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}

func (a *App) GetHistoryByCategoriesVideos(id int, date_start string) (*entity.ProcentByCategoresUser, error) {
	var list *entity.ProcentByCategoresUser
	date_end := "2024-05-20"
	list, err := a.storage.GetHistoryByCategoriesVideos(id, date_start, date_end)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}
