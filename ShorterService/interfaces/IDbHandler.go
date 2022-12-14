package interfaces

import "Linky/ShorterService/models"

type IDbHandler interface {
	InsertShort(short models.Short) error
	SelectLink(short string) (models.Short, error)
}
