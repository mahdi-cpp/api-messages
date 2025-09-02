package application

import (
	"context"
	"sync"

	"github.com/mahdi-cpp/iris-tools/image_loader"
)

type MainStorageManager struct {
	mu         sync.RWMutex
	chats      map[string]*ChatManager // Maps chatIDs to their ChatManager
	iconLoader *image_loader.ImageLoader
	ctx        context.Context
}

func NewMainStorageManager() (*MainStorageManager, error) {

	// Handler the manager
	manager := &MainStorageManager{
		chats: make(map[string]*ChatManager),
		ctx:   context.Background(),
	}

	return manager, nil
}
