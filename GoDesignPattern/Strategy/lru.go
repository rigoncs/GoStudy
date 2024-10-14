package strategy

import "fmt"

type Lru struct{}

func (a *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}
