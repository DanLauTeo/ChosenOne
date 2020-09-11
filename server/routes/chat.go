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
	"sort"
	"strconv"

	//"context"
	"io/ioutil"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	//"reflect"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var ListChatRooms = UserDecorate(listChatRooms)

func listChatRooms(w http.ResponseWriter, r *http.Request, user *models.User) {
	ctx := appengine.NewContext(r)

	dsClient := services.Locator.DsClient()

	var chatrooms = make([]models.ChatRoom, len(user.Chatrooms))

	if err := dsClient.GetMulti(ctx, user.Chatrooms, chatrooms); err != nil {
		log.Printf("Failed to retrieve chatrooms from DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(chatrooms)
	if err != nil {
		log.Printf("Failed to convert Chatrooms to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

var GetChatRoom = UserDecorate(getChatRoom)

func getChatRoom(w http.ResponseWriter, r *http.Request, user *models.User) {
	ctx := appengine.NewContext(r)

	dsClient := services.Locator.DsClient()

	paramID := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		log.Printf("Couldn't parse id from request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := datastore.IDKey("ChatRoom", id, nil)

	var chatroom models.ChatRoom

	if err := dsClient.Get(ctx, key, &chatroom); err != nil {
		log.Printf("Failed to retrieve chatroom from DataStore: %v", err)
		if err == datastore.ErrNoSuchEntity {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}

	if !inList(user.ID, chatroom.Participants) {
		log.Printf("User doesn't belong to this chatroom")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	out, err := json.Marshal(chatroom)
	if err != nil {
		log.Printf("Failed to convert ChatRoom to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

var GetChatRoomMessages = UserDecorate(getChatRoomMessages)

func getChatRoomMessages(w http.ResponseWriter, r *http.Request, user *models.User) {
	ctx := appengine.NewContext(r)

	dsClient := services.Locator.DsClient()

	paramID := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		log.Printf("Couldn't parse id from request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := datastore.IDKey("ChatRoom", id, nil)

	var chatroom models.ChatRoom

	if err := dsClient.Get(ctx, key, &chatroom); err != nil {
		log.Printf("Failed to retrieve chatroom from DataStore: %v", err)
		if err == datastore.ErrNoSuchEntity {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}

	if !inList(user.ID, chatroom.Participants) {
		log.Printf("User doesn't belong to this chatroom")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	messageKeys := make([]*datastore.Key, len(chatroom.Messages))

	for i, id := range chatroom.Messages {
		messageKeys[i] = datastore.IDKey("Message", id, nil)
	}

	var messages []models.Message

	if err := dsClient.GetMulti(ctx, messageKeys, &messages); err != nil {
		log.Printf("Failed to retrieve messages from DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Failed to convert messages to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

var CreateChatRoom = UserDecorate(createChatRoom)

func createChatRoom(w http.ResponseWriter, r *http.Request, user *models.User) {
	ctx := appengine.NewContext(r)

	dsClient := services.Locator.DsClient()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestData map[string]string

	if err := json.Unmarshal(body, &requestData); err != nil {
		log.Printf("Cannot parse json body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestedUser := GetUserByID(ctx, dsClient, requestData["requested_user_id"])

	if requestedUser == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	chatroomKey := datastore.IncompleteKey("ChatRoom", nil)
	chatroom := new(models.ChatRoom)
	chatroom.Participants = []string{user.ID, requestedUser.ID}
	sort.Strings(chatroom.Participants)

	var userChatrooms = make([]models.ChatRoom, len(user.Chatrooms))

	if dsClient.GetMulti(ctx, user.Chatrooms, userChatrooms); err != nil {
		log.Printf("Couldn't determine if a similar chatroom exists: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, existingChatroom := range userChatrooms {
		if listsEqual(existingChatroom.Participants, chatroom.Participants) {
			log.Printf("Chatroom with these participants already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	chatroomKey, err = dsClient.Put(ctx, chatroomKey, chatroom)
	if err != nil {
		log.Printf("Cannot create new chat room: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	chatroom.ID = chatroomKey.ID

	user.Chatrooms = append(user.Chatrooms, chatroomKey)

	if _, err = dsClient.Put(ctx, user.Key, user); err != nil {
		log.Printf("Cannot update user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestedUser.Chatrooms = append(requestedUser.Chatrooms, chatroomKey)

	if _, err = dsClient.Put(ctx, requestedUser.Key, requestedUser); err != nil {
		log.Printf("Cannot update user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(chatroom)
	if err != nil {
		log.Printf("Cannot convert Message to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(out))
}

func inList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func listsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
