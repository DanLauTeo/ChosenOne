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
	"localdev/main/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type ngHandler struct {
	staticPath string
	indexPath  string
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Use(userIdMiddleware)
	for _, route := range apiRoutes {
		router.
			PathPrefix("/api/v1/").
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	angular := ngHandler{staticPath: "./frontend/app/dist/app", indexPath: "index.html"}
	router.PathPrefix("/").Handler(angular)

	return router
}

var apiRoutes = Routes{
	Route{
		"Profile",
		"GET",
		"/user/{id}/",
		GetProfile,
	},
	Route{
		"Customise Profile",
		"PATCH",
		"/user/{id}/",
		EditProfile,
	},
	Route{
		"Update Profile Pic",
		"PUT",
		"/user/{id}/profile-image/",
		ProfilePic,
	},
	Route{
		"List user images",
		"GET",
		"/user/{id}/images/",
		GetUserImages,
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
	// Images
	Route{
		"Handle upload",
		"POST",
		"/images/",
		HandleImageUpload,
	},
	Route{
		"Handle image delete",
		"DELETE",
		"/images/{imageID}/",
		HandleImageDelete,
	},
	// Matches
	Route{
		"Get matcher for current user",
		"GET",
		"/matches/",
		GetMatches,
	},
	// ChatRooms
	Route{
		"Get chatroom by id",
		"GET",
		"/chatrooms/{id}/",
		GetChatRoom,
	},
	Route{
		"List user's chatrooms",
		"GET",
		"/chatrooms/",
		ListChatRooms,
	},
	Route{
		"Create a new chatroom",
		"POST",
		"/chatrooms/",
		CreateChatRoom,
	},
	// Messages
	Route{
		"Get messages from a chatroom",
		"GET",
		"/chatrooms/{id}/messages/",
		GetChatRoomMessages,
	},
	Route{
		"Post a message to a chatroom",
		"POST",
		"/chatrooms/{id}/messages/",
		PostMessageInChatRoom,
	},
	Route{
		"Get a message",
		"GET",
		"/messages/{id}/",
		GetMessage,
	},
	Route{
		"Delete a message",
		"DELETE",
		"/messages/{id}/",
		DeleteMessage,
	},
	// Tasks
	Route{
		"Recalculate user matches",
		"GET",
		"/tasks/recalc-user-matches/",
		HandleRecalcUserMatches,
	},
	// Feed
	Route{
		"Get images for feed",
		"GET",
		"/feed-images/",
		GetPhotosForFeed,
	},
}

func (h ngHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Get filepath
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	//Check whether file exists
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		//File does not exist, serve Angular app
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Otherwise, serve file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func userIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		userService := services.Locator.UserService()
		re := regexp.MustCompile(`($|/)user/me(/|^)`)

		if re.MatchString(r.RequestURI) {
			vars := mux.Vars(r)
			vars["id"] = userService.GetCurrentUserID(ctx)
			r = mux.SetURLVars(r, vars)
		}

		next.ServeHTTP(w, r)
	})
}
