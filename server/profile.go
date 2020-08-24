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
    "net/http"
    "cloud.google.com/go/datastore"
    "github.com/gorilla/mux" 
    "context"
    "fmt"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    dsClient, err := datastore.NewClient(ctx, "capstone-2020-dan-lau-teo")
    if err != nil {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`first save error`))
    }
    user := &User{ "Test name", "12345", "img", "This is my bio. It is long.", "Hello", nil, nil,}
    key := datastore.IncompleteKey("User", nil)
    key, err = dsClient.Put(ctx, key, user)
    if err != nil {
        fmt.Println(err)
    }

    vars := mux.Vars(r)
    userId := vars["id"]

	if err != nil {
		w.WriteHeader(http.StatusOK)
        w.Write([]byte(`first load error`))
	}

    //Retrieves User(?)
    k := datastore.NameKey("User", userId, nil)
    
    var e User 
	if err := dsClient.Get(ctx, k, e); err != nil {
		w.WriteHeader(http.StatusOK)
        w.Write([]byte(`second load error`))
	}

	old := e.Name
	e.Name = "Hello World!"

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		// Handle error.
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Name)
	
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    w.Write([]byte(`called edit`))
}

func ProfilePic(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`called pic`))
}