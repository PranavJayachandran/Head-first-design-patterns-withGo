package main

import "fmt"

type iterator interface {
	hasNext() bool
	getNext() interface{}
}
type folder struct {
	folderName string
	files      []*file
	folders    []*folder
}

func (f *folder) print(count int) {
	for range count {
		fmt.Print("\t")
	}
	fmt.Printf("%s\n", f.folderName)
	for _, val := range f.folders {
		val.print(count + 1)
	}
	if len(f.files) > 0 {
		for range count {
			fmt.Print("\t")
		}
		for _, val := range f.files {
			val.print()
		}
		fmt.Println()
	}
}

type file struct {
	fileName string
}

func (f *file) print() {
	fmt.Print(f.fileName + " ")
}

type folderIterator struct {
	index  int
	folder folder
}

func (i *folderIterator) hasNext() bool {
	if i.index < len(i.folder.folders) {
		return true
	}
	return false
}
func (i *folderIterator) getNext() folder {
	i.index += 1
	return *i.folder.folders[i.index-1]
}
func main() {
	folder := &folder{
		folderName: "Folder 1",
		files:      []*file{},
		folders: []*folder{
			{
				folderName: "Folder 1.1",
				files: []*file{
					{
						fileName: "File1.1",
					},
					{
						fileName: "File 1.2",
					},
				},
				folders: []*folder{
					{
						folderName: "Folder 1.2",
						files:      []*file{},
					},
				},
			},
		},
	}

	// folderIterator := &folderIterator{folder: *folder}
	// for folderIterator.hasNext() {
	// 	fmt.Println(folderIterator.getNext().folderName)
	// }

	folder.print(0)
}

// The implementation for the iterator is just done as an example and doesn't show the accurate usages as multiple classses are mot implementing the required interface
// The Composition pattern is used to create the tree like structure for the folder file storage.
