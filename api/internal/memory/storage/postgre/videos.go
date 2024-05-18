package storage

import (
	"database/sql"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) GetStatusVideos() (*entity.StatusVideos, error) {
	var videos entity.StatusVideos
	query := `SELECT id, name FROM public.starus_video`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.StatusVideo{}

		if err := rows.Scan(
			&video.Id,
			&video.Name,
		); err != nil {
			return nil, err
		}

		videos.AgeCategores = append(videos.AgeCategores, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &videos, nil
}

func (s *Storage) GetAllAgeCategoryVideos() (*entity.AgeCategores, error) {
	var videos entity.AgeCategores
	query := `SELECT id, name FROM public.age_categoty`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.AgeCategory{}

		if err := rows.Scan(
			&video.Id,
			&video.Name,
		); err != nil {
			return nil, err
		}

		videos.AgeCategores = append(videos.AgeCategores, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &videos, nil
}

func (s *Storage) ChekVideo(url string) (bool, error) {
	var count int
	query := `SELECT COUNT(id) FROM public.video WHERE url = $1`
	rows := s.ConnectionDB.QueryRow(query, url)

	err := rows.Scan(
		&count,
	)

	if err == sql.ErrNoRows {
		return false, err
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Storage) AddNewVideo(url string, name string) error {
	var eventId int
	query := `insert into video(url, name) values($1, $2, $3) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, url, 1, name)
	err := result.Scan(&eventId)

	s.log.Debug(eventId)
	if err != nil {
		return err
	}

	return nil
}
