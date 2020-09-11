package routes

import (
	"localdev/main/models"
	"localdev/main/services"
	"net/http"
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"cloud.google.com/go/datastore"
)


func GetPhotosForFeed(w http.ResponseWriter, r *http.Request) {
	var response []string
	ctx := appengine.NewContext(r)
	query := datastore.NewQuery("Image")
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
		response = append(([]models.Feed{}), response)
	}
	out, err := json.Marshal(response)
	w.Write([]byte(out))
}

