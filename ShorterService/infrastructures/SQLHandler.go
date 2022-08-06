package infrastructures

import (
	"Linky/ShorterService/models"
	"database/sql"

	"github.com/lib/pq"
)

type SQLHandler struct {
	Conn *sql.DB
}

func (h *SQLHandler) InsertShort(short models.Short) error {
	statement, err := h.Conn.Prepare("INSERT INTO shorts(link, short) VALUES($1, $2)")
	if err != nil {
		return err.(*pq.Error)
	}
	_, err = statement.Exec(short.Link, short.Short)
	if err != nil {
		return err.(*pq.Error)
	}

	return err
}

func (h *SQLHandler) SelectLink(short string) (models.Short, error) {
	row, err := h.Conn.Query("SELECT link FROM shorts WHERE short = $1", short)
	if err != nil {
		return models.Short{}, err.(*pq.Error)
	}

	defer row.Close()
	var shortBuff models.Short

	for row.Next() {
		row.Scan(&shortBuff.Link)
	}

	return shortBuff, nil
}
