package hashing

type hashMapOpenAdressing struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     []*pair
	TOMBSTONE   *pair
}

func newHashMapOpenAdressing() *hashMapOpenAdressing {
	return &hashMapOpenAdressing{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     make([]*pair, 4),
		TOMBSTONE:   &pair{1, "-1"},
	}
}

func (h *hashMapOpenAdressing) hashFunc(key int) int {
	return key % h.capacity
}

func (h *hashMapOpenAdressing) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func (h *hashMapOpenAdressing) findBucket(key int) int {
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

func (h *hashMapOpenAdressing) put(key int, val string) {
	if h.loadFactor() > h.loadThres {
		h.extend()
	}
	index := h.findBucket(key)
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE {
		h.buckets[index] = &pair{key, val}
		h.size++
	} else {
		h.buckets[index].val = val
	}
}

func (h *hashMapOpenAdressing) get(key int) string {
	index := h.findBucket(key)
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE {
		return ""
	}
	return h.buckets[index].val
}

func (h *hashMapOpenAdressing) remove(key int) {
	index := h.findBucket(key)
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE {
		return
	}
	h.buckets[index] = h.TOMBSTONE
	h.size--
}

func (h *hashMapOpenAdressing) extend() {
	tempBuckets := h.buckets
	h.capacity *= h.extendRatio
	h.buckets = make([]*pair, h.capacity)
	h.size = 0
	for _, pair := range tempBuckets {
		if pair != nil && pair != h.TOMBSTONE {
			h.put(pair.key, pair.val)
		}
	}
}
