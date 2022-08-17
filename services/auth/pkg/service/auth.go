package service

type AuthService struct{}

const userName = "user-name"

func NewAuth() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Authorize(username string) bool {
	if username == userName {
		return true
	}
	return false
}
