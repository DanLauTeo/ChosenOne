package storageClient

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type GCStorageClient struct {
	wrapped *storage.Client
}

func NewGCStorageClient(ctx context.Context) *GCStorageClient {

	gcsClient, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}

	return &GCStorageClient{gcsClient}
}

func (client *GCStorageClient) Upload(ctx context.Context, file multipart.File, bucket, object string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.wrapped.Bucket(bucket).Object(object).NewWriter(ctx)

	file.Seek(0, 0)

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return client.GetServingURL(bucket, object), nil
}

func (client *GCStorageClient) Delete(ctx context.Context, bucket, object string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.wrapped.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}
	fmt.Printf("Blob %v deleted.\n", object)
	return nil
}

func (client *GCStorageClient) GetServingURL(bucket, object string) string {
	return fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucket, object)
}
