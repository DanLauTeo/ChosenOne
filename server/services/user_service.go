package services

import "context"

type UserService interface {
	GetCurrentUserID(c context.Context) string
	IsCurrentUserAdmin(c context.Context) bool
	LoginURL(c context.Context, dest string) string
	LogoutURL(c context.Context, dest string) string
}
