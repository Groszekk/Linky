package interfaces

import "Linky/ShorterService/models"

type ILinkRepository interface {
	AddShort(short models.Short) bool
}
