package routes

import (
	"context"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine"
)

type UserBasedHandlerFunc func(http.ResponseWriter, *http.Request, *models.User)

func UserDecorate(handler UserBasedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		dsClient := services.Locator.DsClient()
		userService := services.Locator.UserService()

		userID := userService.GetCurrentUserID(ctx)

		if userID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := GetUserByID(ctx, dsClient, userID)

		if user == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		handler(w, r, user)
	}
}

func GetUserByID(ctx context.Context, dsClient *datastore.Client, id string) *models.User {
	query := datastore.NewQuery("User").Filter("ID=", id)

	var result []models.User

	if _, err := dsClient.GetAll(ctx, query, &result); err != nil {
		log.Printf("Failed to retreive user from datastore: %v", err)
		return nil
	}

	if len(result) == 0 {
		log.Printf("User doesn't exist in datastore")
		return nil
	}

	return &result[0]
}
