package services

import (
	"log"

	"github.com/bhtoan2204/go-practice.git/dto"
	"github.com/bhtoan2204/go-practice.git/models"
	"github.com/bhtoan2204/go-practice.git/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(input dto.SignUpInput) error
	Login(input dto.LoginInput) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) SignUp(input dto.SignUpInput) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	// Create user model
	user := &models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	// Save user to the database
	err = s.userRepo.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func (s *userService) Login(input dto.LoginInput) (*models.User, error) {
	// Fetch the user by username
	user, err := s.userRepo.GetUserByUsername(input.Username)
	if err != nil {
		log.Printf("User not found: %v", err)
		return nil, err
	}

	// Compare the hashed password with the password in the request
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Printf("Invalid password: %v", err)
		return nil, err
	}

	return user, nil
}
