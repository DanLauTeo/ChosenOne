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
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	//"github.com/gorilla/mux"
)

func GetChatRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dsClient := services.Locator.DsClient()

	userService := services.Locator.UserService()

	userID := userService.GetCurrentUserID(ctx)

	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

	chatrooms := make([]models.ChatRoom, len(user.Conversations))

	if err := dsClient.GetMulti(ctx, user.Conversations, chatrooms); err != nil {
		log.Printf("Cannot retrieve chatrooms from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	out, err := json.Marshal(chatrooms)
	if err != nil {
		log.Printf("Cannot convert Chatrooms to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}
