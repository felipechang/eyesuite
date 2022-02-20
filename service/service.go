package service

import (
	"gitlab.com/hardcake/eyesuite/service/storage"
	"gitlab.com/hardcake/eyesuite/service/tesseract"
	"gitlab.com/hardcake/eyesuite/service/token"
)

type Service struct {
	storage.Storage
	tesseract.Tesseract
	token.Token
}

func NewService(s storage.Storage, t tesseract.Tesseract, k token.Token) *Service {
	return &Service{Storage: s, Tesseract: t, Token: k}
}
