package hashing

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

func (a *arrayHashMap) hashFunc(key int) int {
	return key % 100
}

func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "Not Found!"
	}
	return pair.val
}

func (a *arrayHashMap) put(key int, val string) {
	index := a.hashFunc(key)
	pair := &pair{key: key, val: val}
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) pairSet() []*pair {
	res := make([]*pair, 0)
	for _, pair := range a.buckets {
		if pair != nil {
			res = append(res, pair)
		}
	}
	return res
}

func (a *arrayHashMap) keySet() []int {
	res := make([]int, 0)
	for _, pair := range a.buckets {
		if pair != nil {
			res = append(res, pair.key)
		}
	}
	return res
}

func (a *arrayHashMap) valueSet() []string {
	res := make([]string, 0)
	for _, pair := range a.buckets {
		if pair != nil {
			res = append(res, pair.val)
		}
	}
	return res
}
