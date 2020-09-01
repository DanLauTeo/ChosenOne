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

package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

func CheckDatastore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := datastore.NewQuery("User")
	it := services.Locator.DsClient().Run(ctx, query)
	for {
		var user models.User
		_, err := it.Next(&user)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
		fmt.Fprintf(w, "User %q, ID %q\n", user.Name, user.ID)
	}
	w.Write([]byte("done"))
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Invalid request method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid request method"))
		return
	}

	ctx := r.Context()
	dsClient := services.Locator.DsClient()
	vars := mux.Vars(r)
	userID := vars["id"]

	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no ID provided"))
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user models.User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	out, err := json.Marshal(user)
	if err != nil {
		log.Printf("Cannot convert User to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))

}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		log.Printf("Invalid request method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid request method"))
		return
	}

	//Connect to datastore
	ctx := r.Context()
	dsClient := services.Locator.DsClient()

	//Retrieve user from datastore
	vars := mux.Vars(r)
	userID := vars["id"]
	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no ID provided"))
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user models.User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	//Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get patch operations from body
	var bodyMapped []struct {
		Op    string
		Path  string
		Value string
	}
	err = json.Unmarshal(body, &bodyMapped)
	if err != nil {
		log.Printf("Cannot read patch operations: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read through operations and apply valid changes
	for _, v := range bodyMapped {
		if v.Op == "replace" {
			switch v.Path {
			case "/Name":
				user.Name = v.Value
			case "/Bio":
				user.Bio = v.Value
			case "/ProfilePic":
				user.ProfilePic = v.Value
			}
		}
	}

	//Save patched user
	k, err = dsClient.Put(ctx, k, &user)
	if err != nil {
		log.Printf("Cannot save user to DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Return updated user
	out, err := json.Marshal(user)
	if err != nil {
		log.Printf("Cannot convert user to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func ProfilePic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Printf("Invalid request method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid request method"))
		return
	}

	//Connect to datastore
	ctx := r.Context()
	dsClient := services.Locator.DsClient()
	//Retrieve user from datastore
	vars := mux.Vars(r)
	userID := vars["id"]
	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no ID provided"))
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user models.User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	//Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get patch operations from body
	var bodyMapped []struct {
		Op    string
		Path  string
		Value string
	}
	err = json.Unmarshal(body, &bodyMapped)
	if err != nil {
		log.Printf("Cannot read patch operations: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read through operations and apply valid changes
	for _, v := range bodyMapped {
		if v.Op == "replace" {
			switch v.Path {
			case "/Name":
				user.Name = v.Value
			case "/Bio":
				user.Bio = v.Value
			case "/ProfilePic":
				user.ProfilePic = v.Value
			}
		}
	}

	//Save patched user
	k, err = dsClient.Put(ctx, k, &user)
	if err != nil {
		log.Printf("Cannot save user to DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Return updated user
	out, err := json.Marshal(user)
	if err != nil {
		log.Printf("Cannot convert user to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}
