package user

import (
	"errors"
	"unicode"

	"github.com/similadayo/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository *Repository
	logger     *logging.Logger
}

func NewService(repository *Repository, logger *logging.Logger) *Service {
	return &Service{
		Repository: repository,
		logger:     logger,
	}
}

func (s *Service) CreateUser(username, password, email, firstname, lastname string) (User, error) {
	if err := validatePasswordStrength(password); err != nil {
		return User{}, err
	}

	user := User{
		UserName:  username,
		Password:  password,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
	}

	hashedPassword, err := hashedPassword(password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	createdUser, err := s.Repository.Register(user)
	if err != nil {
		return user, err
	}

	return createdUser, nil
}

func (s *Service) GetUser(email, password string) (User, error) {
	user := User{
		Email:    email,
		Password: password,
	}

	user, err := s.Repository.GetUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetUserByID(userID string) (User, error) {
	user, err := s.Repository.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func hashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hash), nil
}

func CompareHashedPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}

	return nil
}

func validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	//At least 1 upper case
	if !containsUpperCase(password) {
		return errors.New("password must contain at least 1 uppercase character")
	}

	// must contain at least 1 lower case
	if !containsLowerCase(password) {
		return errors.New("password must contain at least 1 lowercase character")
	}

	// must contain at least 1 number
	if !containsNumber(password) {
		return errors.New("password must contain at least 1 number")
	}

	// must contain at least 1 special character
	if !containsSpecial(password) {
		return errors.New("password must contain at least 1 special character")
	}

	return nil
}

func containsUpperCase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowerCase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func containsSpecial(s string) bool {
	for _, char := range s {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return true
		}
	}
	return false
}

func contains(char rune, s string) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}
