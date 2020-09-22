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
	"context"
	"encoding/json"
	"fmt"
	"localdev/main/models"
	"localdev/main/services"

	"log"
	"net/http"
	"net/http/httptest"

	//"os"
	"strings"
	"testing"

	"cloud.google.com/go/datastore"
	jsonpatch "github.com/evanphx/json-patch"
)


func createMockUser(ctx context.Context, dsClient *datastore.Client, id string) *models.User {
	user := models.User{nil, "New User", id, "", "", nil}

	key := datastore.NameKey("User", id, nil)

	key, err := dsClient.Put(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to create user in Datastore: %v", err)
		return nil
	}

	user.Key = key

	key, err = dsClient.Put(ctx, user.Key, &user)
	if err != nil {
		log.Printf("Failed to create user in Datastore: %v", err)
		return nil
	}

	return &user
}

func fillDatastore() {
	ctx := context.Background()
	dsClient := services.Locator.DsClient()
	
	user1 := createMockUser(ctx, dsClient, "mock_user_1")
	_, err := dsClient.Put(ctx, user1.Key, user1)
	if err != nil {
		log.Printf("Failed to update key user in Datastore: %v", err)
	}
	log.Println(user1)

	var user models.User
	if err := dsClient.Get(ctx, user1.Key, &user); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
	}

	log.Println(user)
	fillMockDatastore()
}

func fillMockDatastore(){
	ctx := context.Background()
	dsClient := services.Locator.DsClient()
	main := createMockUser(ctx, dsClient, "mock_user_id")
	fmt.Println(main)
}

func emptyDatastore() {
	ctx := context.Background()
	dsClient := services.Locator.DsClient()
	key := datastore.NameKey("User", "mock_user_1", nil)
	if err := dsClient.Delete(ctx, key); err != nil {
		fmt.Println(err)
	}
}


func TestGetProfileValid(t *testing.T) {
	fillDatastore()
	//Test 1: Valid request
	req, err := http.NewRequest("GET", "/user/mock_user_1/", nil)
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
	newKey := datastore.NameKey("User", "mock_user_1", nil)
	u := models.User{newKey, "New User","mock_user_1", "", "", nil}
	// Check the response body is what we expect.
	json, _ := json.Marshal(u)
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
	req, err := http.NewRequest("GET", "/user/7/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
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
		{"op": "replace", "path": "/Bio", "value": "Successful change"}
	]`)
	req, err := http.NewRequest("PATCH", "/user/mock_user_id/", body)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("here")
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	log.Println("here")
	// Check the response body is what we expect.
	newKey := datastore.NameKey("User", "mock_user_id", nil)
	u := models.User{newKey, "Success","mock_user_id", "", "Successful change", nil}
	json, err := json.Marshal(u)
	if err != nil {
		t.Errorf("handler returned unexpected body")
	}
	expected := string(json)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestEditProfileInvalidReplace(t *testing.T) {
	fillDatastore()
	//Test 2: Change something not allowed
	body := strings.NewReader(`[{"op": "replace", "path": "/ID", "value": "Wronggggg"}]`)

	req, err := http.NewRequest("PATCH", "/user/mock_user_id/", body)
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
	newKey := datastore.NameKey("User", "mock_user_id", nil)
	u := models.User{newKey, "New User","mock_user_id", "", "", nil}
	expected, _ := json.Marshal(u)
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

	req, err := http.NewRequest("PATCH", "/user/mock_user_id/", body)
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
	newKey := datastore.NameKey("User", "mock_user_id", nil)
	u := models.User{newKey, "New User","mock_user_id", "", "", nil}
	expected, _ := json.Marshal(u)
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

	req, err := http.NewRequest("PATCH", "/user/mock_user_id/", body)
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
	newKey := datastore.NameKey("User", "mock_user_id", nil)
	u := models.User{newKey, "New User","mock_user_id", "", "Success", nil}
	expected, _ := json.Marshal(u)
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

/*
func TestProfilePicValid(t *testing.T) {
	fillDatastore()
	fmt.Println(file)
	fmt.Println(i.image)

	req, err := http.NewRequest("PUT", "/user/1/profile-image", file)
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
	u := models.User{"User One", "1", i.image, "User One's Bio", "album", nil, nil}
	expected, _ := json.Marshal(u)
	if !jsonpatch.Equal(expected, rr.Body.Bytes()) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestProfilePicInvalidType(t *testing.T) {
	fillDatastore()
	fmt.Println(file)
	fmt.Println(i.image)

	req, err := http.NewRequest("PUT", "/user/1/profile-image", file)
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
	expected := "Invalid image format"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}

func TestProfilePicEmpty(t *testing.T) {
	fillDatastore()
	req, err := http.NewRequest("PUT", "/user/1/profile-image", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Request body empty"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
	emptyDatastore()
}
*/
