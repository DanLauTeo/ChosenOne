package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"localdev/main/models"
	"localdev/main/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	//"log"

	"cloud.google.com/go/datastore"
)

func fillUserDatastore() {
	createUser("User1","mock_user_id")
	createUser("User2","2")
	createUser("User3","3")
}

func emptyUserDatastore() {
	ctx := context.Background()
	dsClient := services.Locator.DsClient()

	deleteUser("mock_user_id")
	deleteUser("2")
	deleteUser("3")

	var chatrooms []*models.ChatRoom
	keys, err := dsClient.GetAll(ctx, datastore.NewQuery("ChatRoom"), &chatrooms)
	if err != nil {
		fmt.Println(err)
	}
	if err := dsClient.DeleteMulti(ctx, keys); err != nil {
		fmt.Println(err)
	}

	var messages []*models.Message
	keys, err = dsClient.GetAll(ctx, datastore.NewQuery("Message"), &messages)
	if err != nil {
		fmt.Println(err)
	}
	if err := dsClient.DeleteMulti(ctx, keys); err != nil {
		fmt.Println(err)
	}
}

func createUser(name string, id string){
	ctx := context.Background()
	dsClient := services.Locator.DsClient()
	u := models.User{name, id, "img", "Bio", "album", nil, nil}
	k := datastore.NameKey("User", u.ID, nil)
	k, err := dsClient.Put(ctx, k, &u)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteUser(id string){
	ctx := context.Background()
	dsClient := services.Locator.DsClient()
	k := datastore.NameKey("User", id, nil)
	if err := dsClient.Delete(ctx, k); err != nil {
		fmt.Println(err)
	}
}

func TestGetNilChatRooms(t *testing.T) {
	fillUserDatastore()
	//Test 1: Valid request: no chatrooms
	req, err := http.NewRequest("GET", "/messages", nil)
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
	array := []models.ChatRoom{}
	json, _ := json.Marshal(array)
	expected := string(json)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	emptyUserDatastore()
}

func TestCreateAndGetOneChatRoom(t *testing.T) {
	//Test 2: Create One Room and get One Room 
	body := strings.NewReader("2")
	fillUserDatastore()
	req, err := http.NewRequest("POST", "/messages", body)
		if err != nil {
		fmt.Println(err)
	}
	rr1 := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr1, req)
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	req, err = http.NewRequest("GET", "/messages", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := httptest.NewRecorder()
	handler = NewRouter()
	handler.ServeHTTP(rr2, req)
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	if rr1.Body.String() != rr2.Body.String() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr1.Body.String() , rr2.Body.String())
	}
	emptyUserDatastore()
}


func TestPostMessagesValid(t *testing.T) {
	//Test 3: Create A message and post it
	body := strings.NewReader("2")
	fillUserDatastore()
	
	req, err := http.NewRequest("POST", "/messages", body)
		if err != nil {
		fmt.Println(err)
	}
	rr0 := httptest.NewRecorder()
	handler := NewRouter()
	handler.ServeHTTP(rr0, req)
	if status := rr0.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	newChatroom := new(models.ChatRoom)
	
	err = json.Unmarshal([]byte(rr0.Body.String()), &newChatroom)
	if err != nil {
		fmt.Println(err.Error()) 
		//json: Unmarshal(non-pointer main.Request)
	}
	chatroomID := newChatroom.Address

	message := models.Message{"mock_user_id", "Hello", time.Now()}
	m, _ := json.Marshal(message)

	body = strings.NewReader(string(m))

	req, err = http.NewRequest("POST", "/messages/" + chatroomID, body)
		if err != nil {
		fmt.Println(err)
	}
	rr1 := httptest.NewRecorder()
	handler = NewRouter()
	handler.ServeHTTP(rr1, req)
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v. ChatroomID was %v",
			status, http.StatusOK, chatroomID)
	}

	// Check the response body is what we expect.
	req, err = http.NewRequest("GET", "/messages/" + chatroomID, body)
	if err != nil {
	fmt.Println(err)
	}
	rr2 := httptest.NewRecorder()
	handler = NewRouter()
	handler.ServeHTTP(rr2, req)
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}

	// Check the response body is what we expect.

	if rr1.Body.String() != rr2.Body.String() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr1.Body.String(), rr2.Body.String())
	}
	emptyUserDatastore()
}


