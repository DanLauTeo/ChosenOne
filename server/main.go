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

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"localdev/main/dsclient"
	"localdev/main/routes"

	"google.golang.org/appengine"
)

func main() {
	ctx := context.Background()
	dsclient.Init(ctx)
	router := routes.NewRouter()

	if _, local := os.LookupEnv("LOCAL_TESTING"); !local {
		http.Handle("/", router)
		appengine.Main()
	} else {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}

		log.Printf("Listening on port %s", port)

		http.ListenAndServe(":"+port, router)
	}
}
