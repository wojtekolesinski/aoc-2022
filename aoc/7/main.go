package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	s "strings"

	"github.com/wojtekolesinski/aoc-2022/util"
)

//go:embed input.txt
var input string

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

type Dir struct {
	name   string
	parent *Dir
	files  []*File
	dirs   []*Dir
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:   name,
		parent: parent,
		files:  []*File{},
		dirs:   []*Dir{},
	}
}

func (d *Dir) AddFile(file *File) {
	d.files = append(d.files, file)
}

func (d *Dir) AddDir(dir *Dir) {
	d.dirs = append(d.dirs, dir)
}

func (d *Dir) Print(indentLevel ...int) {
	indent := 0
	if len(indentLevel) > 0 {
		indent = indentLevel[0]
	}

	PrintIndent(indent)
	fmt.Printf("- %s (dir)\n", d.name)
	for _, dir := range d.dirs {
		dir.Print(indent + 1)
	}
	for _, file := range d.files {
		PrintIndent(indent + 1)
		fmt.Printf("- %s (file, size=%d)\n", file.name, file.size)
	}
}

func (d *Dir) getSize() int {
	size := 0
	for _, file := range d.files {
		size += file.size;
	}
	for _, dir := range d.dirs {
		size += dir.getSize()
	}
	return size
}

func PrintIndent(howMany int) {
	for i := 0; i < howMany; i++ {
		fmt.Print("  ")
	}
}

func getListOfDirectories(d *Dir) []*Dir {
	dirs := d.dirs

	if len(dirs) == 0 {
		return dirs
	}

	for _, dir := range d.dirs {
		dirs = append(getListOfDirectories(dir), dirs...)
	}
	return dirs
}

func main() {
	fmt.Println(part1())
}

func part1() int {
	lines := s.Split(input, "\n")
	var root *Dir = NewDir("/", nil)
	var currentDir = root
	loop := true
	for index := 1; index < len(lines) && loop; {
		words := s.Split(lines[index], " ")
		switch words[0] {
		case "$":
			handleCommand(words[1:], lines, &index, &currentDir)
		default:
			root.Print()
			loop = false
		}
		
		// fmt.Println(lines[index])
		// root.Print()
		// fmt.Println(root, currentDir)

	}
	size := 0
		for _, dir := range getListOfDirectories(root) {
			dirSize := dir.getSize()
			if dirSize < 100000 {
				size += dirSize
			}
		}
		return size
}

func handleCommand(cmd, lines []string, currentIndex *int, currentDir **Dir) {
	*currentIndex++
	switch cmd[0] {
	case "ls":
		for {
			line := lines[*currentIndex]
			if line[0] == '$' {
				return
			}

			words := s.Split(lines[*currentIndex], " ")
			if words[0] == "dir" {
				dirName := words[1]
				(*currentDir).AddDir(NewDir(dirName, *currentDir))
				// fmt.Println("added dir ", dirName, " to dir ", currentDir.name)
			} else {
				fileName := words[1]
				fileSize, err := strconv.Atoi(words[0])
				util.CheckError(err)

				file := NewFile(fileName, fileSize)
				(*currentDir).AddFile(file)
				// fmt.Println("added file ", *file, " to dir ", currentDir.name)
			}
			*currentIndex++
		}
	case "cd":
		dirName := cmd[1]
		if dirName == ".." {
			// fmt.Println("changing directory to ", (*currentDir).parent.name)
			*currentDir = (*currentDir).parent
			// fmt.Println("current directory: ", (*currentDir).name)

			return
		} else {
			for _, dir := range (*currentDir).dirs {
				if dir.name == dirName {
					// fmt.Println("changing directory to ", dir.name)
					*currentDir = dir
					// fmt.Println("current directory: ", (*currentDir).name)
					return
				}
			}
			log.Fatal("Something went wrong")
		}
	default:
		log.Fatal("unknown command: ", cmd)
	}
}
