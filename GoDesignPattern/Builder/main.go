package builder

import "fmt"

func main() {
	normalBuilder := newNormalBuilder()
	iglooBuilder := newIglooBuilder()
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal house: %s, %s\n", normalHouse.windowType, normalHouse.doorType)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("Igloo house: %s, %s\n", iglooHouse.windowType, iglooBuilder.doorType)
}
