package composite

import "fmt"

type File struct {
	name string
}

func (f *File) find(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
