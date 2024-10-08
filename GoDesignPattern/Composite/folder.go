package composite

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) find(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, v := range f.components {
		v.find(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
