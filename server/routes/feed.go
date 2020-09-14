package routes

import (
	"localdev/main/models"
	"localdev/main/services"
	"net/http"
	"encoding/json"
	"log"

	"google.golang.org/appengine"
	"google.golang.org/api/iterator"
	"cloud.google.com/go/datastore"
)


func GetPhotosForFeed(w http.ResponseWriter, r *http.Request) {
	var response []string
	ctx := appengine.NewContext(r)
	query := datastore.NewQuery("Image").Order("-created").Limit(100) //sorts Images entities by creation time in descending order
	it := services.Locator.DsClient().Run(ctx, query)
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
		feed.ImageURL = image.GCSObjectID()
		k := datastore.NameKey("User", feed.OwnerID, nil)

		var user models.User
		dsClient := services.Locator.DsClient()
		if err = dsClient.Get(ctx, k, &user); err != nil {
			log.Printf("Cannot retrieve user from DataStore: %v", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("entity not found"))
			return
		}
	
		//Get profile pic
		feed.ProfilePic = user.ProfilePic
	
		response = append(response, feed)
	}

	out, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error cannot convert feed to JSON: %v", err)
	}
	w.Write([]byte(out))
}

