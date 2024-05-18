package app

import "github.com/SashaMelva/smart_filter/internal/entity"

func (a *App) GetListChildren(id int) (*entity.ChilgrenLists, error) {
	var list *entity.ChilgrenLists
	list, err := a.storage.GetChildrens(id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}

func (a *App) AddListChildren(idParent, idChil int) error {
	err := a.storage.CreateChildren(idParent, idChil)

	if err != nil {
		a.log.Error(err)
		return err
	}

	return nil
}

func (a *App) GetFiltersByChaild(accoountId int) (*entity.Fileters, error) {
	var filters *entity.Fileters
	filters, err := a.storage.GetFiltersByChaild(accoountId)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return filters, nil
}
