package prototype

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	folder1 := &folder{
		name:  "folder1",
		files: []Inode{file1, file2},
	}

	folder1.print(" ")
	cloneFolder := folder1.clone()
	cloneFolder.print(" ")
}
