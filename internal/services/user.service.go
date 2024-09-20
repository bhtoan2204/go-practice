package services

import "github.com/gin-gonic/gin"

type UserService interface {
	SignUp()
}

type userService struct {
}

// SignUp and Login methods do not return anything, so you should remove the return statements
func (s *userService) SignUp() gin.HandlerFunc {
	// implementation here
	return nil
}
