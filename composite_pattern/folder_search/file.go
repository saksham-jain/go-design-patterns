package foldersearch

import "fmt"

type File struct {
	name string
}

func NewFile(name string) Component {
	return &File{name: name}
}

func (f *File) Search(name string) {
	fmt.Printf("searching file %s for %s\n", f.name, name)
}
