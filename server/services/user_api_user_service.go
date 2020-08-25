package services

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/appengine/user"
)

type UserAPIUserService struct{}

func (_ *UserAPIUserService) GetCurrentUserID(c context.Context) string {
	return user.Current(c).ID
}

func (_ *UserAPIUserService) IsCurrentUserAdmin(c context.Context) bool {
	return user.IsAdmin(c)
}

func (_ *UserAPIUserService) LoginURL(c context.Context, dest string) string {
	url, err := user.LoginURL(c, dest)

	userAPIError(err)

	return url
}

func (_ *UserAPIUserService) LogoutURL(c context.Context, dest string) string {
	url, err := user.LogoutURL(c, dest)

	userAPIError(err)

	return url
}

func userAPIError(err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("User API failed with error: %s", err))
	}
}
