package routes

import (
	"encoding/json"
	"localdev/main/config"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
)

func GetPhotosForFeed(w http.ResponseWriter, r *http.Request) {
	var response []models.Feed
	ctx := appengine.NewContext(r)
	query := datastore.NewQuery("Image").
		Order("-created").
		Filter("type=", "user_uploaded_image").
		Limit(100) //sorts Images entities by creation time in descending order
	dsClient := services.Locator.DsClient()
	storageClient := services.Locator.StorageClient()
	it := dsClient.Run(ctx, query)
	for {
		var feed models.Feed
		var image models.Image
		_, err := it.Next(&image)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error fetching next image: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		feed.OwnerID = image.OwnerID
		feed.ImageURL = storageClient.GetServingURL(config.ImageBucket(), image.GCSObjectID())
		k := datastore.NameKey("User", feed.OwnerID, nil)

		var user models.User
		if err = dsClient.Get(ctx, k, &user); err != nil {
			log.Printf("Cannot retrieve user from DataStore: %v", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("entity not found"))
			return
		}

		//Get profile pic and name
		feed.ProfilePic = user.ProfilePic
		feed.OwnerName = user.Name

		response = append(response, feed)
	}

	out, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error cannot convert feed to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot convert feed to JSON"))
	}
	w.Write([]byte(out))
}
