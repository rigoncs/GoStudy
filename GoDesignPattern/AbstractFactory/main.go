package abstractfactory

import "fmt"

func main() {
	f, _ := GetSportsFactory("nike")
	shoe := f.makeShoe()
	fmt.Printf("Shoe: %s\n", shoe.getLogo())
	shirt := f.makeShirt()
	fmt.Printf("Shirt: %s\n", shirt.getLogo())
}
