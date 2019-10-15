# cacher is a definition with cache

you can use `Register(Cacher)` in your init to implements a `Cacher` then use the interface{}
```go
/*Cacher define an cache interface */
type Cacher interface {
	//get a value
	Get(key string) ([]byte, error)
	//get a value from cache with default
	GetD(key string, v []byte) []byte
	//set a value to cache
	Set(key string, val []byte) error
	//set a value with ttl
	SetWithTTL(key string, val []byte, ttl int64) error
	//check the value is exist
	Has(key string) (bool, error)
	//delete a value
	Delete(key string) error
	//clear all the values
	Clear() error
	//get multi values
	GetMultiple(keys ...string) (map[string][]byte, error)
	//set multi values with [key]value
	SetMultiple(values map[string][]byte) error
	//delete multi values
	DeleteMultiple(keys ...string) error
}
```
