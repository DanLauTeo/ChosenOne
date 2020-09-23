package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"localdev/main/config"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"

	"cloud.google.com/go/datastore"
)

func serveError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Internal Server Error")
	log.Printf("Internal server error: %v", err)
}

func HandleImageUpload(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	userService := services.Locator.UserService()

	userID := userService.GetCurrentUserID(ctx)

	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	file, _, err := r.FormFile("file")

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	defer file.Close()

	visionClient := services.Locator.VisionClient()

	labels, err := visionClient.GetImageLabels(ctx, file)

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	imageURL, imageID, err := UploadImage(ctx, userID, "user_uploaded_image", labels, file)
	if err != nil {
		serveError(ctx, w, err)
		return
	}

	//http.Redirect(w, r, imageURL, http.StatusFound)
	type ImageOut struct {
		ImageID  string `json:"imgID"`
		ImageURL string `json:"imgURL"`
	}

	out, err := json.Marshal(ImageOut{ImageID: fmt.Sprint(imageID), ImageURL: imageURL})
	if err != nil {
		log.Printf("Cannot convert url to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func HandleImageDelete(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	userService := services.Locator.UserService()

	userID := userService.GetCurrentUserID(ctx)

	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Delete
	vars := mux.Vars(r)
	imageID := vars["imageID"]

	err := DeleteImage(ctx, imageID, userID)
	if err != nil {
		serveError(ctx, w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteImage(ctx context.Context, imageID, userID string) error {
	bucket := config.ImageBucket()
	storageClient := services.Locator.StorageClient()
	dsClient := services.Locator.DsClient()

	//Delete from datastore
	//Get key from body
	n, err := strconv.ParseInt(imageID, 10, 64)
	key := datastore.IDKey("Image", n, nil)

	//Get image from datastore
	var image models.Image
	if err := dsClient.Get(ctx, key, &image); err != nil {
		return err
	}

	//Check owner IDs match
	if image.OwnerID != userID {
		err := errors.New("User is not image owner")
		return err
	}

	//Delete
	if err := dsClient.Delete(ctx, key); err != nil {
		return err
	}

	//Delete from GCS
	imageID = image.GCSObjectID()
	err = storageClient.Delete(ctx, bucket, imageID)
	if err != nil {
		return err
	}
	return nil
}

func UploadImage(ctx context.Context, user_id, image_type string, labels []models.LabelWithScore, file multipart.File) (string, int64, error) {
	dsClient := services.Locator.DsClient()

	key := datastore.IncompleteKey("Image", nil)
	entity := new(models.Image)
	entity.Type = image_type
	entity.OwnerID = user_id
	entity.Labels = labels
	entity.Created = time.Now().Unix()

	var err error

	entity.Key, err = dsClient.Put(ctx, key, entity)

	if err != nil {
		return "", 0, err
	}

	bucket := config.ImageBucket()
	object := entity.GCSObjectID()

	storageClient := services.Locator.StorageClient()

	imageURL, err := storageClient.Upload(ctx, file, bucket, object)
	if err != nil {
		return "", 0, err
	}

	return imageURL, entity.Key.ID, nil
}

func GetUserImages(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()
	storageClient := services.Locator.StorageClient()
	vars := mux.Vars(r)
	userID := vars["id"]

	type ImageOut struct {
		ImageID  string `json:"imgID"`
		ImageURL string `json:"imgURL"`
	}

	images := make([]ImageOut, 0)

	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no ID provided"))
		return
	}

	query := datastore.NewQuery("Image").
		Order("-created").
		Filter("owner_id=", userID).
		Filter("type=", "user_uploaded_image")

	it := dsClient.Run(ctx, query)

	for {
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
		images = append(images, ImageOut{ImageID: fmt.Sprint(image.Key.ID), ImageURL: storageClient.GetServingURL(config.ImageBucket(), image.GCSObjectID())})
	}

	out, err := json.Marshal(images)
	if err != nil {
		log.Printf("Cannot convert Images to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}
