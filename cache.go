package cacher

var cache Cacher

/*RegisterCache register cache to map */
func Register(c Cacher) {
	cache = c
}

//Get get value
func Get(key string) ([]byte, error) {
	return cache.Get(key)
}

//GetD get value ,if not found return a default value
func GetD(key string, v []byte) []byte {
	return cache.GetD(key, v)
}

//Set set value
func Set(key string, val []byte) error {
	return cache.Set(key, val)
}

//SetWithTTL set value with time to life
func SetWithTTL(key string, val []byte, ttl int64) error {
	return cache.SetWithTTL(key, val, ttl)
}

//Has check value
func Has(key string) (bool, error) {
	return cache.Has(key)
}

//Delete delete value
func Delete(key string) error {
	return cache.Delete(key)
}

//Clear clear all
func Clear() error {
	return cache.Clear()
}

//GetMultiple get multiple value
func GetMultiple(keys ...string) (map[string][]byte, error) {
	return cache.GetMultiple(keys...)
}

//SetMultiple set multiple value
func SetMultiple(values map[string][]byte) error {
	return cache.SetMultiple(values)
}

//DeleteMultiple delete multiple value
func DeleteMultiple(keys ...string) error {
	return cache.DeleteMultiple(keys...)
}
