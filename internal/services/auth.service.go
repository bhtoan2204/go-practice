package services

type AuthService interface {
	Login()
	Refresh()
}

type authService struct {
}

func NewAuthService() AuthService {
	return nil
}

func (s *authService) Login() {
	// implementation here
}

func (s *authService) Refresh() {
	// implementation here
}
