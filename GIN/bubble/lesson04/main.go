package lesson04

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	name := "Rigon"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v\n", err)
		return
	}
}
