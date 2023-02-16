package randmap

import (
	"math/rand"
	"sync"
)

type element struct {
	content     interface{}
	slice_index int
}

type RandMap struct {
	m sync.RWMutex
	// Where the objects you care about are stored.
	container map[any]element
	// A slice of the map keys used in the map above. We put them in a slice
	// so that we can get a random key by choosing a random index.
	keys []any
	// We store the index of each key, so that when we remove an item, we can
	// quickly remove it from the slice above.
	sliceKeyIndex map[any]int
	//len of keys
	counter int
}

func NewRandMap() *RandMap {
	return &RandMap{
		container:     make(map[any]element),
		sliceKeyIndex: make(map[any]int),
		counter:       0,
	}
}

func (s *RandMap) Count() int {
	return s.counter
}

func (s *RandMap) Set(key any, item interface{}) {
	s.m.Lock()
	defer s.m.Unlock()

	if old_ele, ok := s.container[key]; ok {
		//old exist already
		s.container[key] = element{item, old_ele.slice_index}
	} else {
		// add map key to slice of map keys
		s.keys = append(s.keys, key)
		// store object in map
		s.container[key] = element{item, s.counter}
		s.sliceKeyIndex[key] = s.counter
		s.counter++
	}
}

func (s *RandMap) Get(key any) interface{} {
	s.m.RLock()
	defer s.m.RUnlock()

	if ele, ok := s.container[key]; ok {
		return ele.content
	} else {
		return nil
	}

}

func (s *RandMap) Remove(key any) {
	s.m.Lock()
	defer s.m.Unlock()
	// get index in key slice for key
	index, exists := s.sliceKeyIndex[key]
	if !exists {
		// item does not exist
		return
	}
	delete(s.sliceKeyIndex, key)

	counter_prev := s.counter - 1

	// remove key from slice of keys
	s.keys[index] = s.keys[counter_prev]
	s.keys = s.keys[:counter_prev]

	// we just swapped the last element to another position.
	// so we need to update it's index (if it was not in last position)
	if counter_prev != index { //not the last index
		otherKey := s.keys[index]
		s.sliceKeyIndex[otherKey] = index
	}

	// remove object from map
	delete(s.container, key)

	s.counter--
}

func (s *RandMap) Random() interface{} {

	if s.counter <= 0 {
		return nil
	}

	s.m.RLock()
	defer s.m.RUnlock()

	randomIndex := rand.Intn(s.counter)
	key := s.keys[randomIndex]

	if ele, ok := s.container[key]; ok {
		return ele.content
	} else {
		return nil
	}
}

func (s *RandMap) PopRandom() interface{} {

	if s.counter <= 0 {
		return nil
	}

	s.m.RLock()
	randomIndex := rand.Intn(s.counter)
	key := s.keys[randomIndex]

	item := s.container[key]
	s.m.RUnlock()

	s.Remove(key)

	return item.content
}

func (s *RandMap) Loop(callback func(key,value interface{}) bool) {
	s.m.Lock()
	defer s.m.Unlock()

	for k,v:=range s.container{
		if !callback(k,v){
			return
		}
	}
}
