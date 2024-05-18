package storage

import "time"

func (s *Storage) AddHistory(videoId, children int, date time.Time) error {
	query := `insert into history (account_id, date_time, video_id) values($1, $2)`
	_, err := s.ConnectionDB.Exec(query, children, date, videoId)

	if err != nil {
		return err
	}

	return nil
}
