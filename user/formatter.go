package user

//FormatUser struct for response user
type FormatUser struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

//FormatterUser function for response user
func FormatterUser(user User, token string) FormatUser {
	userformat := FormatUser{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return userformat
}
