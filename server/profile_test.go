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
	"testing"
	"net/http/httptest"
	"net/http"
	"context"
	"cloud.google.com/go/datastore"
	"fmt"
	"encoding/json"
)

func fillDatastore() {
	users := make([]User, 0)
	u := User{"User One", "1", "img", "User One's Bio", "album", nil, nil,}
	users = append(users, u)
	u = User{"User Two", "2", "img", "User Two's Bio", "album", nil, nil,}
	users = append(users, u)
	u = User{"User Three", "3", "img", "User Three's Bio", "album", nil, nil,}
	users = append(users, u)

	for _, user := range users {
		NewProfile(user)
	}
}

func NewProfile(u User) {
    ctx := context.Background()
    dsClient, err := datastore.NewClient(ctx, "lauraod-step-2020")
    if err != nil {
        fmt.Println(err)
    }
    key := datastore.NameKey("User", u.ID,  nil)
    key, err = dsClient.Put(ctx, key, &u)
    if err != nil {
        fmt.Println(err)
    }
}


func TestGetProfile(t *testing.T) {
	fillDatastore()
	//Test 1: Valid request
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	u := User{"User One", "1", "img", "User One's Bio", "album", nil, nil,}
	// Check the response body is what we expect.
	json, _ := json.MarshalIndent(u, "", "  ")
	expected := string(json)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	
	//Test 2: Invalid ID
	req, err = http.NewRequest("GET", "/user/7", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected = `entity not found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEditProfile(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/user/id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `called edit`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestProfilePic(t *testing.T) {
    req, err := http.NewRequest("PUT", "/user/id/profile-image", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

	// Check the response body is what we expect.
	expected := `called pic`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}