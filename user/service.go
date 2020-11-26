package user

import (
	"golang.org/x/crypto/bcrypt"
)

//ServiceInterface ...
type ServiceInterface interface {
	RegisterUser(input RegisterUserInput) (User, error)
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

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return User{}, err
	}

	user := User{
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
