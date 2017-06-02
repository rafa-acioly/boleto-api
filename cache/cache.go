package cache

import (
	"sync"
)

var memcache map[string]interface{}
var mutex sync.Mutex

//Inicia o cache em memoria antes de executar a main
func init() {
	memcache = make(map[string]interface{})
}

//Get item do cache
func Get(key string) (interface{}, bool) {
	k, ok := memcache[key]
	return k, ok
}

//Set item no cache
func Set(key string, obj interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	memcache[key] = obj
}
