package models

import (
	"fmt"

	"cloud.google.com/go/datastore"
)

type Image struct {
	Key     *datastore.Key   `datastore:"__key__"`
	Type    string           `datastore:"type"`
	OwnerID string           `datastore:"owner_id"`
	Labels  []LabelWithScore `datastore:"labels"`
	Created int64            `datastore:"created"`
}

type LabelWithScore struct {
	Name  string  `datastore:"name"`
	Score float32 `datastore:"score"`
}

func (image *Image) GCSObjectID() string {
	if image.Key.ID == 0 {
		return "image_" + image.Key.Name
	}

	return "image_" + fmt.Sprint(image.Key.ID)
}
