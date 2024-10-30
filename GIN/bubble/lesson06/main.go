package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func praise(w http.ResponseWriter, r *http.Request) {
	p := func(name string) (string, error) {
		return name + "年轻又帅气!", nil
	}

	t := template.New("demo.tmpl")
	t.Funcs(template.FuncMap{
		"praise": p,
	})
	_, err := t.ParseFiles("./demo.tmpl")
	if err != nil {
		fmt.Printf("Create template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", praise)
	http.HandleFunc("/tmplDemo", tmplDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, aerr:%v", err)
		return
	}
}
