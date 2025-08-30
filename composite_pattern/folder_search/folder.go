package foldersearch

import "fmt"

type Folder struct {
	name       string
	components []Component
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func (folder *Folder) Search(name string) {
	fmt.Printf("searching folder %s for %s\n", folder.name, name)
	for _, component := range folder.components {
		component.Search(name)
	}
}

func (folder *Folder) Add(components []Component) {
	folder.components = append(folder.components, components...)
}
