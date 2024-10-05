package prototype

import "fmt"

type folder struct {
	files []Inode
	name  string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, file := range f.files {
		file.print(indentation + indentation)
	}
}

func (f *folder) clone() Inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.files {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.files = tempChildren
	return cloneFolder
}
