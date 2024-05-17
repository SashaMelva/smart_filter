package app

import (
	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (a *App) CreateUser(user *entity.User) (int, error) {
	id, err := a.storage.CreateUser(user)

	if err != nil {
		a.log.Error(err)
	} else {
		a.log.Info("Create event whith id = %v", id)
	}

	return id, err
}

func (a *App) GetUserById(id int) (*entity.User, error) {
	var user *entity.User
	user, err := a.storage.GetUserById(id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return user, nil
}

func (a *App) DeleteUser(id int) error {
	err := a.storage.DeleteUserById(id)

	if err != nil {
		a.log.Error(err)
		return err
	}

	return nil
}

func (a *App) UpdateUser(user *entity.User) error {
	err := a.storage.UpdateUser(user)

	if err != nil {
		a.log.Error(err)
		return err
	}

	return nil
}
