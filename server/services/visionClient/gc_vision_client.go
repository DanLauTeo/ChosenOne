package visionClient

import (
	"context"
	"io"
	"localdev/main/models"
	"log"

	vision "cloud.google.com/go/vision/apiv1"
)

type GCVisionClient struct {
	wrapped *vision.ImageAnnotatorClient
}

func NewGCVisionClient(ctx context.Context) *GCVisionClient {

	visionClient, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create Vision API client: %v", err)
	}

	return &GCVisionClient{visionClient}
}

func (client *GCVisionClient) GetImageLabels(ctx context.Context, imageReader io.Reader) ([]models.LabelWithScore, error) {

	image, err := vision.NewImageFromReader(imageReader)

	if err != nil {
		return nil, err
	}

	labels, err := client.wrapped.DetectLabels(ctx, image, nil, 10)

	if err != nil {
		return nil, err
	}

	labelsWithScore := make([]models.LabelWithScore, len(labels))

	for i, label := range labels {
		labelsWithScore[i] = models.LabelWithScore{label.Description, label.Score}
	}

	return labelsWithScore, nil
}
