package factorymethod

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	awm, _ := getGun("awm")
	printDetails(ak47)
	printDetails(awm)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %s\n", g.getPower())
}
