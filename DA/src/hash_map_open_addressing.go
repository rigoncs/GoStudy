package src

import (
	"fmt"
)

type hashMapOpenAddressing struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     []*pair
	TOMBSTONE   *pair
}

func newHashMapOpenAddressing() *hashMapOpenAddressing {
	return &hashMapOpenAddressing{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     make([]*pair, 4),
		TOMBSTONE:   &pair{-1, "-1"},
	}
}

func (h *hashMapOpenAddressing) hashFunc(key int) int {
	return key % h.capacity
}

func (h *hashMapOpenAddressing) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func (h *hashMapOpenAddressing) findBucket(key int) int {
	index := h.hashFunc(key)
	firstTombstone := -1
	for h.buckets[index] != nil {
		if h.buckets[index].key == key {
			if firstTombstone != -1 {
				h.buckets[firstTombstone] = h.buckets[index]
				h.buckets[index] = h.TOMBSTONE
				return firstTombstone
			}
			return index
		}
		if firstTombstone == -1 && h.buckets[index] == h.TOMBSTONE {
			firstTombstone = index
		}
		index = (index + 1) % h.capacity
	}
	if firstTombstone != -1 {
		return firstTombstone
	}
	return index
}

func (h *hashMapOpenAddressing) get(key int) string {
	index := h.findBucket(key)
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		return h.buckets[index].value
	}
	return ""
}

func (h *hashMapOpenAddressing) put(key int, value string) {
	if h.loadFactor() > h.loadThres {
		h.extend()
	}
	index := h.findBucket(key)
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE {
		h.buckets[index] = &pair{key: key, value: value}
		h.size++
	} else {
		h.buckets[index].value = value
	}
}

func (h *hashMapOpenAddressing) remove(key int) {
	index := h.findBucket(key)
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		h.buckets[index] = h.TOMBSTONE
		h.size--
	}
}

func (h *hashMapOpenAddressing) extend() {
	oldBuckets := h.buckets
	h.capacity *= h.extendRatio
	h.buckets = make([]*pair, h.capacity)
	h.size = 0
	for _, pair := range oldBuckets {
		if pair != nil && pair != h.TOMBSTONE {
			h.put(pair.key, pair.value)
		}
	}
}

func (h *hashMapOpenAddressing) print() {
	for _, pair := range h.buckets {
		if pair == nil {
			fmt.Println("nil")
		} else if pair == h.TOMBSTONE {
			fmt.Println("TOMBSTONE")
		} else {
			fmt.Printf("%d -> %s\n", pair.key, pair.value)
		}
	}
}
