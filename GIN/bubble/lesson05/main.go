package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./demo.tmpl")
	if err != nil {
		fmt.Printf("Create template failed, err:%v\n", err)
		return
	}
	user := User{
		Name:   "Rigon",
		Gender: "Male",
		Age:    24,
	}
	m := map[string]interface{}{
		"Name":   "Jeremy",
		"Gender": "Male",
		"Age":    25,
	}
	hobby := []string{
		"乒乓球",
		"羽毛球",
		"双色球",
	}
	tmpl.Execute(w, map[string]interface{}{
		"user":  user,
		"m":     m,
		"hobby": hobby,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, aerr:%v", err)
		return
	}
}
