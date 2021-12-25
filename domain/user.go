package domain

import (
	"time"
)

// User Represents our users in the DataBase
type User struct {
	ID                   string    `json:"id"  san:"trim,xss"`
	Name                 string    `json:"name,omitempty" validate:"my-custom-tag" form:"fname" san:"max=25,trim,title,xss"` // omitempty: if empty, don't encode it at all
	Surname              string    `json:"surname,omitempty" form:"fsurname" san:"max=45,trim,title,xss"`
	IdentificationNumber string    `json:"identification,omitempty" form:"fid" san:"max=45,trim,title,xss"`
	Email                string    `form:"femail" san:"max=50,trim,lower,xss" `
	Password             string    `form:"fpassword" san:"max=50,trim,xss"`
	Address              string    `json:"address,omitempty" form:"faddress" san:"max=80,trim,title,xss"`
	Province             string    `json:"province,omitempty" form:"fprovince" san:"max=80,trim,title,xss"`
	Country              string    `json:"country,omitempty" form:"fcountry"  san:"max=80,trim,xss"`
	State                string    `json:"state,omitempty" form:"fstate" san:"max=80,trim,title,xss"`
	Region               string    `json:"region,omitempty" form:"fregion" san:"max=80,trim,title,xss"`
	CountryPhonePrefix   string    `json:"country-phone-prefix,omitempty" form:"fprefix" san:"max=80,trim,title,xss"`
	Phone                int       `json:"phone,omitempty" form:"fphone" san:"trim,xss"`
	PhotoUser            string    `json:"photo-user,omitempty" san:"xss"`
	PhotoIDFront         string    `json:"photo-id-front,omitempty" san:"xss"`
	PhotoIDBack          string    `json:"photo-id-back,omitempty" san:"xss"`
	PhotoWithID          string    `json:"photo-with-id,omitempty" san:"xss"`
	IsValidated          bool      `json:"validated,omitempty" san:"max=80,trim,title,xss"`
	IsEmailValidated     bool      `json:"emailvalidated,omitempty" san:"max=80,trim,title,xss"`
	Role                 UserRole  `json:"role,omitempty" san:"max=80,trim,title,xss"`
	DateInscription      time.Time `json:"date"`
	FavouriteProperties  []string  `json:"favproperties"  san:"trim,xss"`
}

// UserRole is the close list of different roles that the users can have
type UserRole string

// The options for RoleType
const (
	Customer UserRole = "customer"
	Agent    UserRole = "agent"
	Admin    UserRole = "admin"
)

func (u User) IsAdmin() bool {
	return u.Role == Admin
}

func IsAdmin(u User) bool {
	return u.Role == Admin
}
func IsAgent(u User) bool {
	return u.Role == Agent
}
