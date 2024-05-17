package app

import (
	"errors"

	"github.com/SashaMelva/smart_filter/internal/entity"
	"github.com/SashaMelva/smart_filter/pkg"
)

func (a *App) CreateAccount(account *entity.Account) (int, error) {
	id, err := a.storage.CreateAccount(account)

	if err != nil {
		a.log.Error(err)
	} else {
		a.log.Info("Create event whith id = %v", id)
	}

	return id, err
}

func (a *App) Auth(account *entity.Account) (*entity.Tokens, error) {

	accountId, err := a.storage.GetAccountId(account)
	token, err := a.storage.GetTokenByUser(accountId.Id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}
	if token != nil {
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
		RefreshToken: refToken,
	}

	err = a.storage.SaveTokens(&entity.RefreshToken{
		UserId:       accountId.Id,
		RefreshToken: refToken,
	})

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return &tokens, nil
}
