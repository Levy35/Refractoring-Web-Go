package postgresql

import (
	"database/sql"

	"levijames.net/test/pkg/models"
)

type MusicModel struct {
	DB *sql.DB
}

func (m *MusicModel) Insert(full_name, album, date_released, genre, artist string) (int, error) {
	var id int

	s := `
	INSERT INTO musician(full_name, album, date_released, genre, artist)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING music_id
	`
	err := m.DB.QueryRow(s, full_name, album, date_released, genre, artist).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MusicModel) Read() ([]*models.Music, error) {
	//sql statement
	s := `
	SELECT full_name, album, date_released, genre, artist
	FROM musician
	LIMIT 10
	`

	rows, err := m.DB.Query(s)
	if err != nil {
		return nil, err
	}
	//clean up
	defer rows.Close()

	music := []*models.Music{}

	for rows.Next() {
		m := &models.Music{}

		err = rows.Scan(&m.Full_name, &m.Album, &m.Date_released, &m.Genre, &m.Artist)

		if err != nil {
			return nil, err
		}
		music = append(music, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return music, nil
}
