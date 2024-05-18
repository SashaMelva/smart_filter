package app

import (
	"errors"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (a *App) ChekVideo(video entity.VideoCheker) error {
	ok, err := a.storage.ChekVideo(video.UrlVideo)

	if err != nil {
		a.log.Error(err)
		return err
	}

	if !ok {
		return errors.New("Видео с этим url не существует")
	}

	err = a.storage.AddNewVideo(video.UrlVideo, "")

	if err != nil {
		a.log.Error(err)
	}

	return nil
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
