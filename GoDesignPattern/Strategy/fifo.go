package strategy

import "fmt"

type Fifo struct{}

func (a *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}
