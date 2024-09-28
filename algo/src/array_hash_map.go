package src

import "fmt"

type pair struct {
	key   int
	value string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	buckets := make([]*pair, 100)
	return &arrayHashMap{
		buckets: buckets,
	}
}

func (m *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

func (m *arrayHashMap) get(key int) string {
	index := m.hashFunc(key)
	pair := m.buckets[index]
	if pair == nil {
		return "Not found"
	}
	return pair.value
}

func (m *arrayHashMap) put(key int, val string) {
	pair := &pair{key: key, value: val}
	index := m.hashFunc(key)
	m.buckets[index] = pair
}

func (m *arrayHashMap) remove(key int) {
	index := m.hashFunc(key)
	m.buckets[index] = nil
}

func (m *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range m.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (m *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range m.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (m *arrayHashMap) valSet() []string {
	var values []string
	for _, pair := range m.buckets {
		if pair != nil {
			values = append(values, pair.value)
		}
	}
	return values
}

func (m *arrayHashMap) print() {
	for _, pair := range m.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.value)
		}
	}
}
