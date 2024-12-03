package model

import (
	"fmt"
	"regexp"
)

const (
	USERNAME_ACCOUNT_MAX_LEN  = 18
	USERNAME_ACCOUNT_MIN_LEN  = 4
	USERNAME_PASSWORD_MAX_LEN = 16
	USERNAME_PASSWORD_MIN_LEN = 4
)

type User struct {
	ObjectMeta
	Username   string `json:"username,omitempty" gorm:"index;column:username;unique"`
	Password   string `json:"password,omitempty" gorm:"column:password"`
	Email      string `json:"email,omitempty" gorm:"index;column:email"`
	TotalScore int    `json:"totalscore,omitempty" gorm:"index;column:totalScore"`
	Status     int    `json:"status,omitempty" gorm:"column:status"`
}

type UserList struct {
	ListMeta
	Users []User
}

func (u *User) GetDatabaseName() string {
	return "user"
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func (user User) Validate() error {
	if len(user.Username) < USERNAME_ACCOUNT_MIN_LEN {
		return fmt.Errorf("Username too short")
	}

	if len(user.Username) > USERNAME_ACCOUNT_MAX_LEN {
		return fmt.Errorf("Username too long")
	}

	if len(user.Password) < USERNAME_PASSWORD_MIN_LEN {
		return fmt.Errorf("password too short")
	}

	if len(user.Password) > USERNAME_PASSWORD_MAX_LEN {
		return fmt.Errorf("password too long")
	}

	if !IsValidEmail(user.Email) {
		return fmt.Errorf("email format error")
	}

	return nil
}
