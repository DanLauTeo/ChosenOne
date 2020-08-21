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

package main
 
import (
	"log"
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
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
 
    return router
}
 
var routes = Routes{
    Route{
        "Profile",
        "GET",
        "/users/{id}",
        GetProfile,
    },
    Route{
        "Customise Profile",
        "PATCH",
        "/users/{id}",
        EditProfile,
    },
    Route{
        "Update Profile Pic",
        "PUT",
        "/users/{id}/profile-image",
        ProfilePic,
    },
}

func main() {
 
    router := NewRouter()
 
    log.Fatal(http.ListenAndServe(":8000", router))
}