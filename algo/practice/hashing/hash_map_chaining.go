package hashing

type hashMapChaining struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     [][]pair
}

func newHashMapChaining() *hashMapChaining {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}
	return &hashMapChaining{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

func (m *hashMapChaining) hashFunc(key int) int {
	return key % m.capacity
}

func (m *hashMapChaining) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

func (m *hashMapChaining) get(key int) string {
	index := m.hashFunc(key)
	for _, pair := range m.buckets[index] {
		if pair.key == key {
			return pair.val
		}
	}
	return "Not Found!"
}

func (m *hashMapChaining) put(key int, val string) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	index := m.hashFunc(key)
	for i := range m.buckets[index] {
		if m.buckets[index][i].key == key {
			m.buckets[index][i].val = val
			return
		}
	}
	pair := pair{key: key, val: val}
	m.buckets[index] = append(m.buckets[index], pair)
	m.size++
}

func (m *hashMapChaining) remove(key int) {
	index := m.hashFunc(key)
	for i, pair := range m.buckets[index] {
		if pair.key == key {
			m.buckets[index] = append(m.buckets[index][:i], m.buckets[index][i+1:]...)
			m.size--
			break
		}
	}
}

func (m *hashMapChaining) extend() {
	tmpBuckets := make([][]pair, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tmpBuckets[i] = make([]pair, len(m.buckets[i]))
		copy(tmpBuckets[i], m.buckets[i])
	}

	m.capacity *= m.extendRatio
	m.buckets = make([][]pair, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = make([]pair, 0)
	}
	m.size = 0
	for _, pairs := range m.buckets {
		for _, pair := range pairs {
			m.put(pair.key, pair.val)
		}
	}
}
