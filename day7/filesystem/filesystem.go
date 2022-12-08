package filesystem

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name string

	Directories map[string]Directory
	Files       map[string]File

	Size *int
}

type File struct {
	Name string

	Size int
}

func NewDirectory(name string) Directory {
	var directory Directory

	directory.Name = name
	directory.Directories = make(map[string]Directory)
	directory.Files = make(map[string]File)
	directory.Size = new(int)
	*directory.Size = 0

	return directory
}

func (dir *Directory) PrintStrucure(indentation int) {
	fmt.Printf("- %s (dir, size=%d)\n", dir.Name, dir.Size)
	for _, subDirectory := range dir.Directories {
		for i := 0; i < indentation; i++ {
			fmt.Print("  ")
		}
		subDirectory.PrintStrucure(indentation + 1)
	}
	for _, file := range dir.Files {
		for i := 0; i < indentation; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("- %s (file, size=%d)\n", file.Name, file.Size)
	}
}

func (dir *Directory) ComputeDirectorySize() int {
	*(*dir).Size += (*dir).ComputeFileSizes()

	for _, subDirectory := range (*dir).Directories {
		*(*dir).Size += subDirectory.ComputeDirectorySize()
	}
	fmt.Println((*dir).Name, (*dir).Size)
	return *(*dir).Size
}

func (dir *Directory) ComputeFileSizes() int {
	var sum int

	for _, file := range (*dir).Files {
		sum += file.Size
	}

	return sum
}

func (dir *Directory) GetDirectoryByPath(path string) Directory {
	splittedPath := strings.Split(path, "/")
	directory := *dir
	for _, subDir := range splittedPath {
		if subDir != "" {
			directory = directory.Directories[subDir]
		}
	}
	return directory
}

func (dir *Directory) AddFile(file File) {
	(*dir).Files[file.Name] = file
}

func (dir *Directory) AddDirectory(directory Directory) {
	(*dir).Directories[directory.Name] = directory
}
