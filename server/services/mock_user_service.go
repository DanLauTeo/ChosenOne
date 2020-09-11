package services

import "context"

type MockUserService struct{}

func (_ *MockUserService) GetCurrentUserID(c context.Context) string {
	return "2"
}

func (_ *MockUserService) IsCurrentUserAdmin(c context.Context) bool {
	return true
}

func (_ *MockUserService) LoginURL(c context.Context, dest string) string {
	return dest
}

func (_ *MockUserService) LogoutURL(c context.Context, dest string) string {
	return dest
}
