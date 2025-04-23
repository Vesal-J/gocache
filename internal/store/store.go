package store

type CacheObject struct {
	Key       string
	Value     string
	TTL       int
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
