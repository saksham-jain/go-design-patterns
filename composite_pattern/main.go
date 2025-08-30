package main

import foldersearch "go-design-patterns/composite_pattern/folder_search"

func main() {
	file1 := foldersearch.NewFile("File1")
	file2 := foldersearch.NewFile("File2")

	folder1 := foldersearch.NewFolder("Folder1")
	folder2 := foldersearch.NewFolder("Folder2")

	folder1.Add([]foldersearch.Component{folder2})
	folder2.Add([]foldersearch.Component{file1, file2})

	folder1.Search("SJ")
}
