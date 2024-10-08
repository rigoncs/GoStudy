package composite

func main() {
	f1 := &File{"File1"}
	f2 := &File{"File2"}
	f3 := &File{"File3"}

	folder1 := &Folder{
		name: "Folder1",
	}
	folder1.add(f1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(f2)
	folder2.add(f3)
	folder2.add(folder1)
	folder2.find("rose")
}
