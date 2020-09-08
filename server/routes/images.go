package routes

import (
	"context"
	"fmt"
	"localdev/main/config"
	"localdev/main/models"
	"localdev/main/services"
	"net/http"
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

	dsClient := services.Locator.DsClient()

	key := datastore.IncompleteKey("Image", nil)
	entity := new(models.Image)
	entity.Type = "user_uploaded_image"
	entity.OwnerID = userID
	entity.Labels = labels

	entity.Key, err = dsClient.Put(ctx, key, entity)

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	bucket := config.ImageBucket()
	object := entity.GCSObjectID()

	storageClient := services.Locator.StorageClient()

	imageURL, err := storageClient.Upload(ctx, file, bucket, object)
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

	bucket := config.ImageBucket()

	vars := mux.Vars(r)
	image := vars["imageID"]

	storageClient := services.Locator.StorageClient()

	err := storageClient.Delete(ctx, bucket, image)
	if err != nil {
		serveError(ctx, w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}