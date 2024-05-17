package storage

import (
	"context"
	"database/sql"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) SaveTokens(tokens *entity.RefreshToken) error {

	query := `UPDATE token SET token = $1 where account_id = $2`
	_, err := s.ConnectionDB.Exec(query, tokens.RefreshToken, tokens.UserId)

	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) GetTokenByUser(id int) (*entity.RefreshToken, error) {
	var token *entity.RefreshToken
	query := `select token from token where account_id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		token.RefreshToken,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return token, err
	}

	token.UserId = id

	return token, nil
}

func (s *Storage) UpdateTokenByUser(reefToken *entity.RefreshToken, ctx context.Context) error {
	return nil
}
