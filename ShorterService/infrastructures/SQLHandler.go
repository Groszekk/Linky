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
