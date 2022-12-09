package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day7.txt
var input string

type file struct {
	name     string
	children map[string]*file
	parent   *file
	_size    int
	isDir    bool
}

type fileSystem struct {
	root *file
	path map[*file][]*file
}

func newFileSystem() fileSystem {
	dir := newDir("/")
	path := make(map[*file][]*file)

	path[dir] = []*file{}

	return fileSystem{root: dir, path: path}
}

func newFile(name string, size int) *file {
	return &file{name: name, isDir: false, _size: size}
}

func newDir(name string) *file {
	return &file{name: name, isDir: true, children: make(map[string]*file)}
}

func (f *file) size() int {
	if !f.isDir {
		return f._size
	}

	if f._size > 0 {
		return f._size
	}

	for _, child := range f.children {
		f._size += child.size()
	}

	return f._size
}

func (f *file) search(max int) []*file {
	res := []*file{}
	if f.isDir {
		if f.size() < max {
			res = append(res, f)
		}
		for _, child := range f.children {
			res = append(res, child.search(max)...)
		}
	}
	return res
}
func (f *file) dirSize() []int {
	res := []int{}
	if f.isDir {
		res = append(res, f.size())
		for _, child := range f.children {
			res = append(res, child.dirSize()...)
		}
	}
	return res
}

func main() {
	fmt.Println("Total size of the directories is:", part1(input))
	fmt.Println("Smallest directory is:", part2(input))
}

func buildFileSystem(input string) fileSystem {
	var (
		currentDir *file
		dirs       []*file
	)

	fs := newFileSystem()
	instructions := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	currentDir = &file{name: "/", isDir: true, children: make(map[string]*file)}

	for _, instruction := range instructions {
		// A command was recognised
		if strings.HasPrefix(instruction, "$") {
			cmd := strings.TrimPrefix(instruction, "$ ")
			if strings.HasPrefix(cmd, "cd") {
				dir := strings.TrimPrefix(cmd, "cd ")
				switch dir {
				case "/":
					for file := range fs.path {
						if file.name == dir {
							currentDir = file
						}
					}
				case "..":
					currentDir = currentDir.parent
				default:
					// check to see if the directory already exists.
					if _, ok := currentDir.children[dir]; ok {
						currentDir = currentDir.children[dir]
					}
				}
			}
		} else {
			// if the string starts with an integer
			if strings.HasPrefix(instruction, "dir") {
				dir := strings.TrimPrefix(instruction, "dir ")
				if _, ok := currentDir.children[dir]; ok {
					continue
				}

				d := newDir(dir)

				d.parent = currentDir
				currentDir.children[d.name] = d
				fs.path[currentDir] = append(fs.path[currentDir], d)

				dirs = append(dirs, d)

				continue
			} else {
				c := strings.Split(instruction, " ")
				conv, err := strconv.Atoi(c[0])
				if err != nil {
					fmt.Println(err)
				}

				f := newFile(c[1], conv)
				f.parent = currentDir

				currentDir.children[f.name] = f
			}
		}

	}
	return fs
}

func part1(input string) int {
	const fileSizeLimit = 100000
	var total int

	fs := buildFileSystem(input)
	for _, file := range fs.root.search(fileSizeLimit) {
		total += file.size()
	}

	return total
}

func part2(input string) int {
	fs := buildFileSystem(input)

	totalSize := fs.root.size()
	max := 70000000
	needed := 30000000
	unused := max - totalSize

	toDelete := needed - unused
	size := fs.root.dirSize()
	sort.Ints(size)

	i := sort.Search(len(size), func(i int) bool { return size[i] >= toDelete })

	return size[i]
}
