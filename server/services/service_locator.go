package services

import (
	"os"
)

type ServiceLocator struct{}

func (_ *ServiceLocator) UserService() UserService {
	return userService
}

var (
	Locator     ServiceLocator = ServiceLocator{}
	userService UserService
)

func init() {
	_, local := os.LookupEnv("LOCAL_TESTING")

	if local {
		userService = &MockUserService{}
	} else {
		userService = &UserAPIUserService{}
	}
}
