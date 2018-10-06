package auth

import "fmt"

type User struct {
	Id       int
	UserName string
	Password string
}

func (u *User) ToString() string {
	return fmt.Sprintf("id : %d, userName: %s", u.Id, u.UserName)
}
