package store

import "time"

type CacheType string

const (
	STRING CacheType = "string"
	HASH   CacheType = "hash"
	LIST   CacheType = "list"
	SET    CacheType = "set"
	ZSET   CacheType = "zset"
)

type CacheObject struct {
	Type      CacheType
	Value     interface{}
	TTL       time.Duration
	CreatedAt int64
}

type Store struct {
	Caches     map[string]CacheObject
	TTLManager *TTLManager
}

func NewStore() *Store {
	store := &Store{
		Caches: make(map[string]CacheObject),
	}

	store.TTLManager = NewTTLManager(store)
	store.TTLManager.Start()
	return store
}
