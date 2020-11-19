package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// User : Contain infos of a user
type User struct {
	UUID      string                `json:"uuid"`
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	LoginAt   time.Time             `json:"login_at"`
	CreatedAt time.Time             `json:"created_at"`
	Errors    []errconv.PiccoloHccError `json:"errors"`
}

// UserList : Contain list of users
type UserList struct {
	Users  []User                `json:"user_list"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// UserNum : Contain the number of users
type UserNum struct {
	Number int                   `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// Token : Contain the user token
type Token struct {
	Token  string                `json:"token"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

// IsValid : Contain the validation of the token
type IsValid struct {
	IsValid bool                  `json:"isvalid"`
	Errors  []errconv.PiccoloHccError `json:"errors"`
}
