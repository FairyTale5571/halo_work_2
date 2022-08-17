package service

import (
	"time"
)

const (
	signingKey = "vs,.d[/124ecw@#!%^!@v,asd"
	tokenTTL   = 12 * time.Hour
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Authorization interface {
	Authorize(username string) bool
	GenerateToken(username string) (string, error)
	ParseToken(token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{
		Authorization: NewAuth(),
	}
}
