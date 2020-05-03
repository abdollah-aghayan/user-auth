package domain

import (
	"user-auth/utils/errorh"
	"user-auth/utils/validationh"
)

//User struct
type User struct {
	ID       string `json:"id" sql:"id" bson:"_id"`
	Username string `json:"username" sql:"username" bson:"username"`
	Email    string `json:"email" sql:"email" bson:"email"`
	Password string `json:"password,omitempty" sql:"password" bson:"password"`
}

// RegisterUser struct
type RegisterUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

// LoginUser struct
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidateLoginUser check weather requested data for login is ok or not
func (u *LoginUser) ValidateLoginUser() *errorh.Errorh {
	if !validationh.IsValidEmail(u.Email) {
		return errorh.BadRequestError("Invalid email")
	}

	if len(u.Password) == 0 {
		return errorh.BadRequestError("Invalid password")
	}

	return nil
}

// ValidateLoginUser check weather requested data for login is ok or not
func (u *RegisterUser) ValidateRegisterUser() *errorh.Errorh {
	if !validationh.IsValidEmail(u.Email) {
		return errorh.BadRequestError("Invalid email")
	}

	if len(u.Username) == 0 {
		return errorh.BadRequestError("Username can not be empty")
	}

	if len(u.Password) == 0 || len(u.RepeatPassword) == 0 {
		return errorh.BadRequestError("Invalid password")
	}

	if len(u.RepeatPassword) == 0 {
		return errorh.BadRequestError("Invalid repeat password")
	}

	if u.Password != u.RepeatPassword {
		return errorh.BadRequestError("Password should be equal to repeat password")
	}

	return nil
}
