package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name   string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"Golang", 2001, 0, []string{"python, java"}}

	//编码的过程 结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程  json -> 结构体
	movie2 := Movie{}
	err = json.Unmarshal(jsonStr, &movie2)
	if err != nil {
		fmt.Println("json unmarshal error")
		return
	}
	fmt.Printf("%v\n", movie2)
}
