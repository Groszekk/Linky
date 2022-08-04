package repositories

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"fmt"
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
		fmt.Println(err.Error())
		return false
	}

	return true
}
