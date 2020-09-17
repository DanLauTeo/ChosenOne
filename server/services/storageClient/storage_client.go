package storageClient

import (
	"context"
	"mime/multipart"
)

type StorageClient interface {
	Upload(ctx context.Context, file multipart.File, bucket, object string) (string, error)
	Delete(ctx context.Context, bucket, object string) (error)
	GetServingURL(bucket, object string) (string)
}
