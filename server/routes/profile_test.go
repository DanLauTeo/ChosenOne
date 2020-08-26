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
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"localdev/main/config"
	"localdev/main/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func fillDatastore() {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, config.Project())
	if err != nil {
		fmt.Println(err)
	}
	u := models.User{"User One", "1", "img", "User One's Bio", "album", nil, nil}
	key := datastore.NameKey("User", u.ID, nil)
	key, err = dsClient.Put(ctx, key, &u)
	if err != nil {
		fmt.Println(err)
	}
}

func emptyDatastore() {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, config.Project())
	if err != nil {
		fmt.Println(err)
	}
	key := datastore.NameKey("User", "1", nil)
	if err := dsClient.Delete(ctx, key); err != nil {
		fmt.Println(err)
	}
}

func TestGetProfileValid(t *testing.T) {
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
	u := models.User{"User One", "1", "img", "User One's Bio", "album", nil, nil}
	// Check the response body is what we expect.
	json, _ := json.MarshalIndent(u, "", "  ")
	expected := string(json)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	emptyDatastore()
}

func TestGetProfileInvalid(t *testing.T) {
	fillDatastore()
	//Test 2: Invalid ID
	req, err := http.NewRequest("GET", "/user/7", nil)
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
	expected := `entity not found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	emptyDatastore()
}

func TestEditProfileValid(t *testing.T) {
	fillDatastore()
	//Test 1: Change name, description, profile photo
	body := strings.NewReader(`[{"op": "replace", "path": "/Name", "value": "Success"},
		{"op": "replace", "path": "/Bio", "value": "Successful change"},
		{"op": "replace", "path": "/ProfilePic", "value": "Successful change"}
	]`)

	req, err := http.NewRequest("PATCH", "/user/1", body)
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
	u := models.User{"Success", "1", "Successful change", "Successful change", "album", nil, nil}
	expected, _ := json.MarshalIndent(u, "", "  ")
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestEditProfileInvalidReplace(t *testing.T) {
	fillDatastore()
	//Test 2: Change something not allowed
	body := strings.NewReader(`[{"op": "replace", "path": "/ID", "value": "Wronggggg"}]`)

	req, err := http.NewRequest("PATCH", "/user/1", body)
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
	u := models.User{"User One", "1", "img", "User One's Bio", "album", nil, nil}
	expected, _ := json.MarshalIndent(u, "", "  ")
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestEditProfileInvalidOps(t *testing.T) {
	fillDatastore()
	//Test 3: Remove or add a field
	body := strings.NewReader(`[{"op": "remove", "path": "/Name"},
		{"op": "add", "path": "/NewPath", "value": "not allowed"}
	]`)

	req, err := http.NewRequest("PATCH", "/user/1", body)
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
	u := models.User{"User One", "1", "img", "User One's Bio", "album", nil, nil}
	expected, _ := json.MarshalIndent(u, "", "  ")
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestEditProfileMixed(t *testing.T) {
	fillDatastore()
	//Test 4: If receiving a mix of legal and illegal patches, apply legal patches
	body := strings.NewReader(`[{"op": "remove", "path": "/Name"},
		{"op": "replace", "path": "/Bio", "value": "Success"},
		{"op": "replace", "path": "/ID", "value": "Wronggggg"}
	]`)

	req, err := http.NewRequest("PATCH", "/user/1", body)
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
	u := models.User{"User One", "1", "img", "Success", "album", nil, nil}
	expected, _ := json.MarshalIndent(u, "", "  ")
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
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
