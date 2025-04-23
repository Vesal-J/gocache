package store

type CacheObject struct {
	Key   string
	Value string
	TTL   int
}

type Store struct {
	Caches map[string]CacheObject
}

func NewStore() *Store {
	return &Store{
		Caches: make(map[string]CacheObject),
	}
}
