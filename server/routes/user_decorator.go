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

		user, err := GetOrCreateUserByID(ctx, dsClient, userID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		handler(w, r, user)
	}
}

func GetUserByID(ctx context.Context, dsClient *datastore.Client, id string) (*models.User, error) {

	key := datastore.NameKey("User", id, nil)

	var user models.User

	err := dsClient.Get(ctx, key, &user)

	return &user, err
}

func createUserInDatatore(ctx context.Context, dsClient *datastore.Client, id string) (*models.User, error) {
	user := models.User{nil, "New User", id, "", "", nil}

	key := datastore.NameKey("User", id, nil)

	key, err := dsClient.Put(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to create user in Datastore: %v", err)
		return nil, err
	}

	user.Key = key

	return &user, nil
}

func GetOrCreateUserByID(ctx context.Context, dsClient *datastore.Client, id string) (*models.User, error) {
	user, err := GetUserByID(ctx, dsClient, id)

	if err == datastore.ErrNoSuchEntity {
		user, err = createUserInDatatore(ctx, dsClient, id)
	}

	return user, err
}
