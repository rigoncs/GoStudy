package main

import "fmt"

func main1() {
	var map1 map[string]string
	fmt.Printf("map1 is %v\n", map1)
	if map1 == nil {
		fmt.Println("map1 is null")
	}

	map1 = make(map[string]string, 10)
	map1["one"] = "golang"
	map1["two"] = "python"
	fmt.Printf("map1 is %v\n", map1)

	map2 := make(map[int]string)
	map2[1] = "Golang"
	fmt.Printf("map2 is %v\n", map2)

	map3 := map[int]string{
		1: "Csharp",
		2: "Java",
		3: "C++",
	}
	fmt.Printf("map3 is %v\n", map3)

}
