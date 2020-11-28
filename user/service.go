package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//ServiceInterface ...
type ServiceInterface interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
}

//Service ...
type Service struct {
	repo RepositoryInterFace
}

//NewService ..
func NewService(repository RepositoryInterFace) *Service {
	return &Service{repository}
}

//RegisterUser ...
func (s *Service) RegisterUser(input RegisterUserInput) (User, error) {

	user, err := s.repo.FindByEmail(input.Email)

	if user.ID != 0 {
		return user, errors.New("Email already exist")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return User{}, err
	}

	user = User{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   string(hashedPass),
		Role:       "User",
	}

	newUser, err := s.repo.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

//Login ...
func (s *Service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	pass := input.Password

	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if err != nil {
		return user, errors.New("password yang anda berikan salah")
	}

	return user, nil
}
