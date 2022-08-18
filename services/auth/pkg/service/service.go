package service

type Authorization interface {
	Authorize(username string) bool
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{
		Authorization: NewAuth(),
	}
}
