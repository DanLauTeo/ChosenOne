package routes

import (
	"context"
	"errors"
	"fmt"
	"localdev/main/config"
	"localdev/main/models"
	"localdev/main/services"
	"mime/multipart"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"cloud.google.com/go/datastore"
)

func serveError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Internal Server Error")
	log.Errorf(ctx, "%v", err)
}

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

// For temporary purposes, to be deleted later
const rootTemplateHTML = `
	<html>
		<body>
			<form action="{{.}}" method="POST" enctype="multipart/form-data">
				Upload File: <input type="file" name="file">
				<br>
				<input type="submit" name="submit" value="Submit">
			</form>
		</body>
	</html>
	`

func ImageUploadPage(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	uploadURL := "/image-uploaded/"

	w.Header().Set("Content-Type", "text/html")

	if err := rootTemplate.Execute(w, uploadURL); err != nil {
		log.Errorf(ctx, "%v", err)
	}
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

	imageURL, err := UploadImage(ctx, userID, "user_uploaded_image", labels, file)
	if err != nil {
		serveError(ctx, w, err)
		return
	}

	http.Redirect(w, r, imageURL, http.StatusFound)
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
	imageID = "image_" + imageID
	err = storageClient.Delete(ctx, bucket, imageID)
	if err != nil {
		return err
	}
	return nil
}

func UploadImage(ctx context.Context, user_id, image_type string, labels []models.LabelWithScore, file multipart.File) (string, error) {
	dsClient := services.Locator.DsClient()

	key := datastore.IncompleteKey("Image", nil)
	entity := new(models.Image)
	entity.Type = image_type
	entity.OwnerID = user_id
	entity.Labels = labels

	var err error

	entity.Key, err = dsClient.Put(ctx, key, entity)

	if err != nil {
		return "", err
	}

	bucket := config.ImageBucket()
	object := entity.GCSObjectID()

	storageClient := services.Locator.StorageClient()

	imageURL, err := storageClient.Upload(ctx, file, bucket, object)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}
