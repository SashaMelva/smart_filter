package app

import (
	"errors"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/SashaMelva/smart_filter/pkg"
)

func (a *App) CreateAccount(account *entity.Account) (int, error) {
	ok, err := a.storage.ChekLoginAccount(account)

	if err != nil {
		a.log.Error(err)
		return 0, err
	}
	if ok != 0 {
		err = errors.New("Пользователь с таким логином уже существует")
		a.log.Error(err)
		return 0, err
	}

	id, err := a.storage.CreateAccount(account)

	if err != nil {
		a.log.Error(err)
		return 0, err
	} else {
		a.log.Info("Create event whith id = %v", id)
	}

	err = a.storage.CreateToken(id)

	if err != nil {
		a.log.Error(err)
		return 0, err
	}

	return id, err
}

func (a *App) Auth(account *entity.Account) (*entity.Tokens, error) {

	accountId, err := a.storage.GetAccountId(account)
	a.log.Debug(accountId)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	token, err := a.storage.GetTokenByUser(accountId.Id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}
	if token != nil && token.RefreshToken != "" && token.RefreshToken != " " {
		a.log.Debug(token)
		return nil, errors.New("В базе уже существует токен для данного пользователя")
	}

	accessToken, err := pkg.GenerateAccssesToken(accountId.Id, a.Tokens.SecretJWT)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	refToken, err := pkg.GenerateRefreshToken()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	tokens := entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: string(refToken),
	}

	err = a.storage.SaveTokens(&entity.RefreshToken{
		UserId:       accountId.Id,
		RefreshToken: accessToken,
	})

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return &tokens, nil
}
