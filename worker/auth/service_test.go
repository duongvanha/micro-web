package auth

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetAccountWrongPath(t *testing.T) {

	Convey("Test auth service", t, func() {

		userName := "userName"
		password := "password"

		service := Service{}

		Convey("Test auth service", func() {

			Convey("When user if not found", func() {

				Convey("Login should return error", func() {
					_, e := service.Login(userName, password)
					So(e, ShouldBeError, errors.New(userName+" not found"))
				})

				Convey("Register should return nil", func() {
					e := service.Register(userName, password)
					So(e, ShouldBeNil)
				})

			})

			Convey("When users are available", func() {

				service.Register(userName, password)

				Convey("Login should return user and error nil", func() {
					user, err := service.Login(userName, password)
					So(err, ShouldBeNil)

					So(user, ShouldNotBeNil)
				})

				Convey("Register should return error", func() {
					e := service.Register(userName, password)
					So(e, ShouldBeError, errors.New("the "+userName+" is already is used"))
				})

			})
		})
	})
}
