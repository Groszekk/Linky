package services

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type LinkService struct {
	interfaces.ILinkRepository
}

func (s *LinkService) Short(short models.Short) (bool, string) {
	randStr := RandString(4)
	short.Short = randStr

	success := s.AddShort(short)
	if !success {
		return false, *new(string)
	}

	return true, randStr
}

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
