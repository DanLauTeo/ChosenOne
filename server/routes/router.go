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

/*This file manages the different routes of the project*/

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range apiRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var apiRoutes = Routes{
	Route{
		"Profile",
		"GET",
		"/user/{id}",
		GetProfile,
	},
	Route{
		"Customise Profile",
		"PATCH",
		"/user/{id}",
		EditProfile,
	},
	Route{
		"Update Profile Pic",
		"PUT",
		"/user/{id}/profile-image",
		ProfilePic,
	},
	Route{
		"Test Profile",
		"GET",
		"/",
		CheckDatastore,
	},
	Route{
		"Get login URL",
		"GET",
		"/login-url",
		LoginURL,
	},
	Route{
		"Get logout URL",
		"GET",
		"/logout-url",
		LogoutURL,
	},
	Route{
		"Get current user id",
		"GET",
		"/who",
		Who,
	},
	Route{
		"Upload page",
		"GET",
		"/upload/",
		ImageUploadPage,
	},
	Route{
		"Handle upload",
		"POST",
		"/image-uploaded/",
		HandleImageUpload,
	},
	Route{
		"Get chatrooms from user",
		"GET",
		"/messages",
		GetChatRooms,
	},
	Route{
		"Get messages from chatroom",
		"GET",
		"/messages/{chatroomID}",
		GetMessagesFromChatRoom,
	},
}
