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
	//"context"
	"io/ioutil"
	"localdev/main/models"
	"localdev/main/services"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
)

func CheckChatDatastore(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := datastore.NewQuery("Chat")
	it := services.Locator.DsClient().Run(ctx, query)
	for {
		var chatroom models.ChatRoom
		_, err := it.Next(&chatroom)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error fetching next chatroom: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ChatRoom %q, Messages %q\n", chatroom.ID, chatroom.Messages)
	}
	w.Write([]byte("done"))
}

func GetMessagesFromChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()
	vars := mux.Vars(r)
	chatroomID := vars["chatroomID"]

	userService := services.Locator.UserService()
	userID := userService.GetCurrentUserID(ctx)

	k := datastore.NameKey("ChatRoom", chatroomID, nil)

	var chatroom models.ChatRoom

	if err := dsClient.Get(ctx, k, &chatroom); err != nil {
		log.Printf("Cannot retrieve chatroom from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	if userInList(userID, chatroom.Participants) {
		log.Printf("User not participant of chatroom")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User not participant of chatroom"))
		return
	}

	messages := make([]models.Message, len(chatroom.Messages)) //Will be max 100 messages
	err := dsClient.GetMulti(ctx, chatroom.Messages, messages)
	if err != nil {
		log.Printf("Cannot get messages: %v", err)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(messages)
	if err != nil {
		log.Printf("Cannot convert Message to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetChatRooms(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()

	userService := services.Locator.UserService()

	userID := userService.GetCurrentUserID(ctx)

	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	k := datastore.NameKey("User", userID, nil)

	var user models.User
	if err := dsClient.Get(ctx, k, &user); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("entity not found"))
		return
	}

	chatrooms := make([]models.ChatRoom, len(user.Chatrooms))

	if err := dsClient.GetMulti(ctx, user.Chatrooms, chatrooms); err != nil {
		log.Printf("Cannot retrieve chatrooms from DataStore: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
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

func userInList(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()
	userService := services.Locator.UserService()
	userID := userService.GetCurrentUserID(ctx)

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	requestedUserID := string(bodyBytes)
	id := userID + " " + requestedUserID
	participants := []string{userID, requestedUserID}
	cr := models.ChatRoom{ id, participants, nil}
	
	chatroomKey := datastore.NameKey("ChatRoom", cr.ID, nil)
	_, err1 := dsClient.Put(ctx, chatroomKey, &cr)
	if err1 != nil {
		fmt.Println(err)
	}
	
	//add chatroomKey to CurrentUser.ChatRooms
	k1 := datastore.NameKey("User", userID, nil)
	
	var currentUser models.User
	if err := dsClient.Get(ctx, k1, &currentUser); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		return
	}
	
	currentUser.Chatrooms = append(currentUser.Chatrooms, chatroomKey)
	
	k1, err = dsClient.Put(ctx, k1, &currentUser)
	if err != nil {
		log.Printf("Cannot save user to DataStore: %v", err)
		return
	}
	// add chatroomid to RequestedUser.ChatRooms
	k2 := datastore.NameKey("User", requestedUserID, nil)
	
	var requestedUser models.User
	if err := dsClient.Get(ctx, k2, &requestedUser); err != nil {
		log.Printf("Cannot retrieve user from DataStore: %v", err)
		return
	}
	
	requestedUser.Chatrooms = append(requestedUser.Chatrooms, chatroomKey)
	
	k2, err = dsClient.Put(ctx, k2, &currentUser)
	if err != nil {
		log.Printf("Cannot save user to DataStore: %v", err)
		return
	}
}

func PostMessageChatRoom(w http.ResponseWriter, r *http.Request){  
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()
	vars := mux.Vars(r)
	chatroomID := vars["chatroomID"]

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	var message models.Message
	err = json.Unmarshal(bodyBytes, &message)
	if err != nil {
		log.Printf("Cannot parse message to json body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Setting the key name to be "<time><userID>"
	keyname := message.Posted.String() + message.SenderID
	newKey := datastore.NameKey("Message", keyname, nil)
	newKey, err = dsClient.Put(ctx, newKey, &message)
	if err != nil {
		fmt.Println(err)
	}
	k := datastore.NameKey("ChatRoom", chatroomID, nil)

	var chatroom models.ChatRoom

	if err := dsClient.Get(ctx, k, &chatroom); err != nil {
		log.Printf("Cannot retrieve chatroom from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}
	
	if len(chatroom.Messages) > 100 {
		//Delete first message in  list
		keyOfFirstMessage := chatroom.Messages[0]
		if err := dsClient.Delete(ctx, keyOfFirstMessage); err != nil {
			log.Printf("Cannot delete first message: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

 }
 
 func DeleteMessageChatRoom(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	dsClient := services.Locator.DsClient()
	vars := mux.Vars(r)
	chatroomID := vars["chatroomID"]

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var message models.Message
	err = json.Unmarshal(bodyBytes, &message)
	if err != nil {
		log.Printf("Cannot parse message to json body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get messages search for the one that is meant to be deleted, and delete it
	k := message.Posted.String() + message.SenderID
	keyOfUnwantedMessage := datastore.NameKey("Message", k, nil)
	if err := dsClient.Delete(ctx, keyOfUnwantedMessage); err != nil {
		log.Printf("Cannot delete first message: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Delete the key from chatroom.Messages
	key := datastore.NameKey("ChatRoom", chatroomID, nil)

	var chatroom models.ChatRoom

	if err := dsClient.Get(ctx, key, &chatroom); err != nil {
		log.Printf("Cannot retrieve chatroom from DataStore: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("entity not found"))
		return
	}

	for i := len(chatroom.Messages) - 1; i >= 0; i-- {
		if chatroom.Messages[i] == keyOfUnwantedMessage {
			chatroom.Messages = append(chatroom.Messages[:i], chatroom.Messages[i+1:]...)
			break
		}
	}
	
	// Updating chatroom messages
	key, err = dsClient.Put(ctx, key, &chatroom)
	if err != nil {
		log.Printf("Cannot save chatroom to DataStore: %v", err)
		return
	}
}

