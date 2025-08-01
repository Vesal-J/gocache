package store

import (
	"sync"
	"time"
)

type TTLManager struct {
	store    *Store
	mutex    sync.RWMutex
	stopChan chan struct{}
}

func NewTTLManager(store *Store) *TTLManager {
	return &TTLManager{
		store:    store,
		stopChan: make(chan struct{}),
	}
}

func (tm *TTLManager) Start() {
	go tm.cleanupLoop()
}

func (tm *TTLManager) Stop() {
	close(tm.stopChan)
}

func (tm *TTLManager) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tm.cleanup()
		case <-tm.stopChan:
			return
		}
	}
}

func (tm *TTLManager) cleanup() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	now := time.Now().Unix()
	for key, obj := range tm.store.Caches {
		if obj.TTL > 0 && obj.CreatedAt+int64(obj.TTL.Seconds()) < now {
			delete(tm.store.Caches, key)
		}
	}
}
