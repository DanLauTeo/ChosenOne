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

package models

import (
	"time"

	"cloud.google.com/go/datastore"
)

type Message struct {
	ID        int64     `json:"id" datastore:"-"`
	SenderID  string    `json:"sender_id"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
}

func (x *Message) LoadKey(k *datastore.Key) error {
	x.ID = k.ID
	return nil
}

func (x *Message) Load(ps []datastore.Property) error {
	// Load as usual.
	return datastore.LoadStruct(x, ps)
}

func (x *Message) Save() ([]datastore.Property, error) {
	// Save as usual.
	return datastore.SaveStruct(x)
}
