package main

type KeyValue struct {
	store map[string]string
}

func (kv *KeyValue) Set(key string, val string) {
	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) string {
	return kv.store[key]
}

func NewStorage() *KeyValue {
	return &KeyValue{
		store: make(map[string]string),
	}
}
