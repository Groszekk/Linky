package infrastructures

import (
	"Linky/ShorterService/models"
	"database/sql"
)

type SQLHandler struct {
	Conn *sql.DB
}

func (h *SQLHandler) InsertShort(short models.Short) error {
	statement, err := h.Conn.Prepare("INSERT INTO shorts(link, short) VALUES($1, $2)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(short.Link, short.Short)
	if err != nil {
		return err
	}

	return err
}
