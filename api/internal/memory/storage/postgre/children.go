package storage

import (
	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) CreateChildren(children *entity.Chilgren) (int, error) {
	var childrenId int
	query := `insert into children(user_parent_id, user_id) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, children.IdParent, children.IdUser)
	err := result.Scan(&childrenId)

	if err != nil {
		return 0, err
	}

	return childrenId, nil
}

func (s *Storage) DeleteChildren(childrenId int) error {
	query := `delete from children where user_id = $1`
	_, err := s.ConnectionDB.Exec(query, childrenId)

	if err != nil {
		return err
	}

	return nil
}
