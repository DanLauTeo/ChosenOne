// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dsclient

import (
	"context"
	"localdev/main/config"
	"log"

	"cloud.google.com/go/datastore"
)

var dsClient *datastore.Client

func Init(ctx context.Context) {
	var err error
	dsClient, err = datastore.NewClient(ctx, config.Project())
	if err != nil {
		log.Fatalf("Cannot connect to DataStore: %v", err)
	}
}

func DsClient() *datastore.Client {
	if dsClient != nil {
		return dsClient
	} else {
		log.Fatal("Datastore client not initialised")
		return nil
	}
}
