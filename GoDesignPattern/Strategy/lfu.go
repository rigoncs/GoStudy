package strategy

import "fmt"

type Lfu struct{}

func (a *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}
