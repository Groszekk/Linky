package interfaces

import "Linky/ShorterService/models"

type ILinkService interface {
	Short(short models.Short) (bool, string)
}
