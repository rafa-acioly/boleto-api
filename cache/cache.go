package cache

var memcache map[string]interface{}

//Inicia o cache em memoria antes de executar a main
func init() {
	memcache = make(map[string]interface{})
}

//Get item do cache
func Get(key string) interface{} {
	return memcache[key]
}

//Set item no cache
func Set(key string, obj interface{}) {
	memcache[key] = obj
}
