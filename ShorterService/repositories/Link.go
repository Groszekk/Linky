package repositories

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"log"
)

type LinkRepositoryWithCircuitBreaker struct {
	interfaces.ILinkRepository
}

type LinkRepository struct {
	interfaces.IDbHandler
}

func (r *LinkRepository) AddShort(short models.Short) bool {
	err := r.InsertShort(short)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}

	return true
}
