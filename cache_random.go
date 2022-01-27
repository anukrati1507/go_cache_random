package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Data struct {
	data           int64
	expirationTime int64
	index          int
}
type InMap struct {
	Map map[string]*Data
}
type Parameters struct {
	data int64
	key  string
	ttl  int64
}

func (c *InMap) set(param Parameters) (item *Data, exists bool) {
	if param.ttl == 0 {
		param.ttl = 60
	} else if param.ttl < 0 {
		panic("Time cannot be a negative entity")
	}
	_, exists = c.Map[param.key]
	c.Map[param.key] = &Data{data: param.data, expirationTime: time.Now().Unix() + param.ttl}
	item = c.Map[param.key]
	return
}
func (c *InMap) get(key string) (value int64) {
	var state bool
	var item *Data

	if time.Now().Unix() > c.Map[key].expirationTime {
		delete(c.Map, key)
	}
	item, state = c.Map[key]
	value = item.data
	if !state {
		panic("Key not found")
	}
	return
}
func (c *InMap) delete(key string) {
	var state bool
	_, state = c.Map[key]
	if !state {
		panic("Key not found")
	}
	delete(c.Map, key)

	return
}
func (c *InMap) deleteRandom(key string) {
	var state bool
	_, state = c.Map[key]
	if state {
		//panic("Key not found")
		if time.Now().Unix() > c.Map[key].expirationTime {
			delete(c.Map, key)
		}
	}
	//delete(c.Map, key)

	return
}
func (c *InMap) printMap() {
	for key, value := range c.Map {
		fmt.Printf("%s: %d %d \n", key, value.data, value.expirationTime)
	}
}
func MapRandomKeyGet(mapI *InMap) {
	keys := reflect.ValueOf(mapI).MapKeys()
	keydel := keys[rand.Intn(len(keys))].Interface()
	mapI.deleteRandom(keydel.(string))
	return
}
func checkExpiry(inmap *InMap) {
	for x := range time.Tick(5 * time.Minute) {
		//f(x)'
		fmt.Println(x)
		for i := 1; i < 10; i++ {
			if len(inmap.Map) > 10 {
				MapRandomKeyGet(inmap)
			}
		}
	}
}
func newCache() (inmap InMap) {
	inmap = InMap{Map: make(map[string]*Data)}
	return
}
