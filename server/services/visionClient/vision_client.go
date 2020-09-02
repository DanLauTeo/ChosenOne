package visionClient

import (
	"context"
	"io"
	"localdev/main/models"
)

type VisionClient interface {
	GetImageLabels(ctx context.Context, imageReader io.Reader) ([]models.LabelWithScore, error)
}
