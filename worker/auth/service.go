package auth

import "errors"

type Service struct {
	users []User
}

func (auth *Service) Login(userName string, password string) (User, error) {

	for _, user := range auth.users {
		if user.UserName == userName && user.Password == password {
			return user, nil
		}
	}

	return User{}, errors.New(userName + " not found")
}

func (auth *Service) Register(userName string, password string) error {
	for _, user := range auth.users {
		if user.UserName == userName && user.Password == password {
			return errors.New("the " + userName + " is already is used")
		}
	}
	users := append(auth.users, User{UserName: userName, Password: password})

	auth.users = users

	return nil
}
