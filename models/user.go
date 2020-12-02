package models

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" mapstructure:"last_name"`
	IsActive  bool   `json:"is_active"`
}

type PublicUser struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) Public() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func PublicUsers(users []User) []PublicUser {
	var publicUsers []PublicUser
	for _, u := range users {
		publicUsers = append(publicUsers, u.Public())
	}
	return publicUsers
}

func (u *User) Init(email, firstName, lastName string) User {
	return User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (u *User) Setter(newKey uint) {
	u.ID = newKey
}

func (u User) Getter() uint {
	return u.ID
}

func (u User) TableName() string {
	return "users"
}
