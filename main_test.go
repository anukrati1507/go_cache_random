package main

import (
	"testing"
	"time"
)

func TestSetValue(t *testing.T) {
	cache := newCache()
	go checkExpiry(&cache)
	cache.set(Parameters{key: "X", data: 100, ttl: 4})
	time.Sleep(2 * time.Second)
	cache.set(Parameters{key: "Y", data: 34, ttl: 4})
	time.Sleep(2 * time.Second)
	cache.get("X")
	time.Sleep(2 * time.Second)
	cache.get("Y")
	time.Sleep(2 * time.Second)
	if len(cache.Map) != 0 {
		cache.printMap()
		t.Errorf("Map is not empty")
	}
}

func TestResetValue(t *testing.T) {
	cache := newCache()
	go checkExpiry(&cache)
	var data int64
	cache.set(Parameters{key: "X", data: 100, ttl: 100})
	time.Sleep(2 * time.Second)
	cache.set(Parameters{key: "X", data: 34, ttl: 100})
	data = cache.get("X")

	if data != 34 {
		t.Errorf("Reset not worked")
	}
}

func TestDeleteOperation(t *testing.T) {
	cache := newCache()
	go checkExpiry(&cache)
	cache.set(Parameters{key: "X", data: 50, ttl: 10})
	cache.set(Parameters{key: "Y", data: 29, ttl: 89})
	cache.delete("X")
	if len(cache.Map) != 1 {
		cache.printMap()
		t.Errorf("Expected 1 element in map")
	}
	cache.delete("Y")
	if len(cache.Map) != 0 {
		t.Errorf("Expected zero elements in map")
	}

}

func GetValueOperation(t *testing.T) {
	cache := newCache()
	go checkExpiry(&cache)
	cache.set(Parameters{key: "X", data: 50, ttl: 5})
	cache.set(Parameters{key: "Y", data: 29, ttl: 10})
	time.Sleep(5 * time.Second)
	cache.get("X")
	if len(cache.Map) != 1 {
		t.Errorf("Deletion of X not successful even on time limit exceed")
	}

}
