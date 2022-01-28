package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//for storing key value and expiration time
type Data struct {
	data           int64
	expirationTime int64
	index          int
}

// creating a map
type InMap struct {
	Map map[string]*Data
}

//input arguments
type Parameters struct {
	data int64
	key  string
	ttl  int64
}

//set function for cache
func (c *InMap) set(param Parameters) (item *Data, exists bool) {
	if param.ttl == 0 {
		param.ttl = 60 // default value if time not given
	} else if param.ttl < 0 {
		panic("Time cannot be a negative entity") //error if time is negative
	}
	_, exists = c.Map[param.key] //check if the value exists already or not
	c.Map[param.key] = &Data{data: param.data, expirationTime: time.Now().Unix() + param.ttl}
	item = c.Map[param.key] //store value in the map
	return
}

//get function to get value for provided key value
func (c *InMap) get(key string) (value int64) {
	var state bool
	var item *Data
	item, state = c.Map[key]
	if !state {
		panic("Key not found")
	}

	if time.Now().Unix() > c.Map[key].expirationTime {
		delete(c.Map, key) //check if expiration time has exceeded and delete the item
		//panic("Key not found")
	} else {
		item, state = c.Map[key]
		value = item.data
		if !state {
			panic("Key not found")
		}
	}
	return
}

//delete function to delete provided key value pair
func (c *InMap) delete(key string) {
	var state bool
	_, state = c.Map[key]
	// if !state {
	// 	panic("Key not found")
	// }
	if state {
		delete(c.Map, key)
		fmt.Println("Key-value pair with key - ", key, ", deleted successfully")
	}
	return
}

//delete randomly generated expired key value pair from the map
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

//print the values present in the map
func (c *InMap) printMap() {
	for key, value := range c.Map {
		fmt.Println(key, value.data, value.expirationTime)
	}
}

//generate random keys for the process of deletion every 20 seconds
func MapRandomKeyGet(mapI *InMap) {
	keys := reflect.ValueOf(mapI).MapKeys()
	keydel := keys[rand.Intn(len(keys))].Interface()
	mapI.deleteRandom(keydel.(string))
	return
}

//check expiry of elements which are not used yet
func checkExpiry(inmap *InMap) {
	for x := range time.Tick(20 * time.Second) {
		//f(x)'
		fmt.Println(x)
		for i := 1; i < 3; i++ {
			if len(inmap.Map) > 3 {
				MapRandomKeyGet(inmap)
			}
		}
	}
}

//create a newcache to store the data
func newCache() (inmap InMap) {
	inmap = InMap{Map: make(map[string]*Data)}
	return
}
