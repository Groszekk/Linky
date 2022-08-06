package interfaces

import "Linky/ShorterService/models"

type ILinkService interface {
	Short(short models.Short) (string, bool)
	GetRedirectLink(shortEndpoint string) (string, bool)
}
