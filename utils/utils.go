package utils

import (
	"errors"

	. "github.com/ninefive/ninefive-blog/lib"
	"github.com/ninefive/ninefive-blog/models"
)

func CheckLogin(username, password string) (user models.User, err error) {
	user = models.GetUserByUsername(username)
	if user.Id == "" {
		return user, errors.New("The user is not exists.")
	}
	if user.Password != Pwshash(password) {
		return user, errors.New("The password is wrong.")
	}
	return user, nil
}
