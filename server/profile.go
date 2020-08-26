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
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"net/http"
)

func CheckDatastore(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, Conf.Project)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("first save error"))
	}
	query := datastore.NewQuery("User")
	it := client.Run(ctx, query)
	for {
		var user User
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
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, Conf.Project)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("datastore error"))
		return
	}
	vars := mux.Vars(r)

	userID := vars["id"]

	if len(userID) == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("no ID provided"))
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("entity not found"))
		return
	}

	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error converting to json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))

}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	//Connect to datastore
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, Conf.Project)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("datastore error"))
		return
	}

	//Retrieve user from datastore
	vars := mux.Vars(r)
	userID := vars["id"]
	if len(userID) == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("no ID provided"))
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("entity not found"))
		return
	}

	//Get request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error reading body"))
		return
	}

	//Get patch operations from body
	var bodyMapped []map[string]string
	err = json.Unmarshal(body, &bodyMapped)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error reading body"))
		return
	}

	//Read through operations and apply valid changes
	for _, v := range bodyMapped {
		if v["op"] == "replace" {
			switch v["path"] {
			case "/Name":
				user.Name = v["value"]
			case "/Bio":
				user.Bio = v["value"]
			case "/ProfilePic":
				user.ProfilePic = v["value"]
			}
		}
	}

	//Save patched user
	k, err = dsClient.Put(ctx, k, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error storing new user"))
		return
	}

	//Return updated user
	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error converting to json"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func ProfilePic(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("called pic"))
}
