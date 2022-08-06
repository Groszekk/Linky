package services

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type LinkService struct {
	interfaces.ILinkRepository
}

func (s *LinkService) Short(short models.Short) (string, bool) {
	randStr := RandString(4)
	short.Short = randStr

	success := s.AddShort(short)
	if !success {
		return *new(string), false
	}

	return randStr, true
}

func (s *LinkService) GetRedirectLink(short string) (string, bool) {
	link, success := s.GetLink(short)
	if !success {
		return *(new(string)), false
	}

	return link, true

}

func RandString(length int) string {
	b := make([]byte, length)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
