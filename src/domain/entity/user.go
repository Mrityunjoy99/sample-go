package entity

import "github.com/Mrityunjoy99/sample-go/src/infrastructure/database"

type User struct {
	database.BaseModel
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func (u User) TableName() string {
	return "users"
}

func NewUser(firstName, lastName, email, phone string) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
}
