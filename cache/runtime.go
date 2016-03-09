package cache

import "github.com/byorty/hardcore/types"

var runtime types.Cache

func Runtime() types.Cache {
	if runtime == nil {
		runtime = make(RuntimeImpl)
	}
	return runtime
}

type RuntimeImpl map[string]interface{}

func (r RuntimeImpl) Get(key string) interface{} {
	if value, ok := r.has(key); ok {
		return value
	} else {
		return nil
	}
}

func (r RuntimeImpl) has(key string) (interface{}, bool) {
	return r[key]
}

func (r RuntimeImpl) Has(key string) bool {
	_, ok := r.has(key)
	return ok
}

func (r *RuntimeImpl) Set(key string, value interface{}) types.Cache {
	r[key] = value
	return r
}
