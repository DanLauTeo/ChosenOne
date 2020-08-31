package services

import (
	"context"
	"localdev/main/services/dsclient"
	"log"
	"os"

	"cloud.google.com/go/datastore"
)

type ServiceLocator struct{}

func (_ *ServiceLocator) UserService() UserService {
	return userService
}

func (_ *ServiceLocator) DsClient() *datastore.Client {
	if dsClient != nil {
		return dsClient
	} else {
		log.Fatal("Datastore client not initialised")
		return nil
	}
}

var (
	Locator     ServiceLocator = ServiceLocator{}
	userService UserService
	dsClient    *datastore.Client
)

func init() {
	_, local := os.LookupEnv("LOCAL_TESTING")

	if local {
		userService = &MockUserService{}
	} else {
		userService = &UserAPIUserService{}
	}

	dsClient = dsclient.NewDatastoreClient(context.Background())
}
