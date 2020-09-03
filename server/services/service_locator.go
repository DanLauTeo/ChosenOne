package services

import (
	"context"
	"localdev/main/services/dsclient"
	sc "localdev/main/services/storageClient"
	vc "localdev/main/services/visionClient"
	"os"

	"cloud.google.com/go/datastore"
)

type ServiceLocator interface {
	UserService() UserService
	DsClient() *datastore.Client
	StorageClient() sc.StorageClient
	VisionClient() vc.VisionClient
}

type DefaultServiceLocator struct{}

func (_ *DefaultServiceLocator) UserService() UserService {
	return userService
}

func (_ *DefaultServiceLocator) DsClient() *datastore.Client {
	return dsClient
}

func (_ *DefaultServiceLocator) StorageClient() sc.StorageClient {
	return storageClient
}

func (_ *DefaultServiceLocator) VisionClient() vc.VisionClient {
	return visionClient
}

var (
	Locator       ServiceLocator = &DefaultServiceLocator{}
	userService   UserService
	dsClient      *datastore.Client
	storageClient sc.StorageClient
	visionClient  vc.VisionClient
)

func init() {
	_, local := os.LookupEnv("LOCAL_TESTING")

	if local {
		userService = &MockUserService{}
	} else {
		userService = &UserAPIUserService{}
	}

	ctx := context.Background()

	dsClient = dsclient.NewDatastoreClient(ctx)

	storageClient = sc.NewGCStorageClient(ctx)

	visionClient = vc.NewGCVisionClient(ctx)
}
