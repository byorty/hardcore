package cache

import (
	"github.com/byorty/hardcore/types"
	"sync"
)

var runtime types.Cache

func Runtime() types.Cache {
	if runtime == nil {
		runtime = &RuntimeImpl{
			mtx:  new(sync.Mutex),
			data: make(map[string]interface{}),
		}
	}
	return runtime
}

type RuntimeImpl struct {
	mtx  *sync.Mutex
	data map[string]interface{}
}

func (r RuntimeImpl) Get(key string) interface{} {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if value, ok := r.get(key); ok {
		return value
	} else {
		return nil
	}
}

func (r RuntimeImpl) get(key string) (interface{}, bool) {
	value, ok := r.data[key]
	return value, ok
}

func (r RuntimeImpl) Has(key string) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	_, ok := r.get(key)
	return ok
}

func (r *RuntimeImpl) Set(key string, value interface{}) types.Cache {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.data[key] = value
	return r
}

func (r *RuntimeImpl) Remove(key string) types.Cache {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.data, key)
	return r
}

func (r RuntimeImpl) GetKeys() []string {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	keys := make([]string, 0)
	for key, _ := range r.data {
		keys = append(keys, key)
	}
	return keys
}
