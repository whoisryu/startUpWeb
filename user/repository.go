package user

import (
	"gorm.io/gorm"
)

//RepositoryInterFace ...
type RepositoryInterFace interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

//Repository ...
type Repository struct {
	db *gorm.DB
}

//NewRepositoryUser ...
func NewRepositoryUser(db *gorm.DB) *Repository {
	return &Repository{db}
}

//Save for insert data user to database
func (r *Repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

//FindByEmail finding user by email
func (r *Repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
