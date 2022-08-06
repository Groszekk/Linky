package repositories

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"fmt"

	"github.com/lib/pq"
	"github.com/omeid/pgerror"
)

type LinkRepositoryWithCircuitBreaker struct {
	interfaces.ILinkRepository
}

type LinkRepository struct {
	interfaces.IDbHandler
}

func (r *LinkRepository) AddShort(short models.Short) bool {
	err := r.InsertShort(short)

	if err, success := err.(*pq.Error); success {
		if e := pgerror.UniqueViolation(err); e != nil {
			_err := r.InsertShort(short)
			if _err != nil {
				fmt.Println(_err.Error())
				return false
			}
			return true
		}
		return false
	}

	return true
}
