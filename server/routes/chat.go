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
	//"io/ioutil"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

func CheckChatDatastore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := datastore.NewQuery("Chat")
	it := services.Locator.DsClient().Run(ctx, query)
	for {
		var chat_room models.ChatRoom
		_, err := it.Next(&chat_room)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
		fmt.Fprintf(w, "Chat %q, ID %q\n", chat_room.Name, chat_room.Messages)
	}
	w.Write([]byte("done"))
}

func GetMessagesFromChatRoom(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		log.Printf("Invalid request method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid request method"))
		return
	}

	ctx := r.Context()
	dsClient := services.Locator.DsClient()
	vars := mux.Vars(r)
	chat_roomID := vars["chat_roomID"]

	k := datastore.NameKey("ChatRoom", chat_roomID, nil)

	var chat_room models.ChatRoom

	if err := dsClient.Get(ctx, k, &chat_room); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	var messages[] models.Message
	err := dsClient.GetMulti(ctx, chat_room.Messages, messages)
	if err != nil {
		log.Printf("Cannot get messages: %v", err)
	}

	out, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Cannot convert Message to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

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