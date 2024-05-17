package storage

import (
	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) CreateChildren(children, idParent int) error {
	var childrenId int
	query := `insert into children(user_parent_id, user_id) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, idParent)
	err := result.Scan(&childrenId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetChildrens(idParent int) (*entity.ChilgrenLists, error) {
	var children entity.ChilgrenLists
	query := `insert into children(user_parent_id, user_id) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, idParent)
	err := result.Scan(&children)

	if err != nil {
		return nil, err
	}

	return &children, nil
}

func (s *Storage) DeleteChildren(childrenId int) error {
	query := `delete from children where user_id = $1`
	_, err := s.ConnectionDB.Exec(query, childrenId)

	if err != nil {
		return err
	}

	return nil
}
