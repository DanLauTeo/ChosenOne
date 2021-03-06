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

package config

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

var config = struct {
	Project        string `json:"project"`
	ImageBucket    string `json:"image_bucket"`
	MatcherAddress string `json:"matcher_address"`
}{}

func Project() string {
	return config.Project
}

func ImageBucket() string {
	return config.ImageBucket
}

func MatcherAddress() string {
	return config.MatcherAddress
}
