package routes

import (
	"context"
	"fmt"
	"io"
	"localdev/main/config"
	"localdev/main/models"
	"localdev/main/services"
	"mime/multipart"
	"net/http"
	"text/template"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
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
	file, _, err := r.FormFile("file")

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	defer file.Close()

	labels, err := getImageLabels(ctx, file)

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	dsClient := services.Locator.DsClient()

	key := datastore.IncompleteKey("Image", nil)
	entity := new(models.Image)
	entity.Type = "user_uploaded_image"
	entity.Labels = labels

	entity.Key, err = dsClient.Put(ctx, key, entity)

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	bucket := config.ImageBucket()
	object := entity.GCSObjectID()

	client, err := storage.NewClient(ctx)

	if err != nil {
		serveError(ctx, w, err)
		return
	}

	defer client.Close()

	if err := gcsUpload(ctx, client, file, bucket, object); err != nil {
		serveError(ctx, w, err)
		return
	}

	imageURL := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucket, object)

	http.Redirect(w, r, imageURL, http.StatusFound)
}

func gcsUpload(ctx context.Context, client *storage.Client, file multipart.File, bucket, object string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)

	file.Seek(0, 0)

	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func getImageLabels(ctx context.Context, file multipart.File) ([]models.LabelWithScore, error) {

	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	image, err := vision.NewImageFromReader(file)

	if err != nil {
		return nil, err
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)

	if err != nil {
		return nil, err
	}

	labelsWithScore := make([]models.LabelWithScore, len(labels))

	for i, label := range labels {
		labelsWithScore[i] = models.LabelWithScore{label.Description, label.Score}
	}

	return labelsWithScore, nil
}
