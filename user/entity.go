package user

import (
	"time"
)

// User ...
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	Password       string
	Role           string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
