package d07

import (
	"strconv"
	"strings"
)

type nodeType int

const (
	file nodeType = iota
	dir
)

type node struct {
	kind     nodeType
	name     string
	size     int
	parent   *node
	children []*node
}

func (n *node) addChildNode(parts []string) {
	// already exists?
	for _, child := range n.children {
		if child.name == parts[1] {
			return
		}
	}

	if parts[0] == "dir" {
		n.children = append(n.children, &node{dir, parts[1], 0, n, make([]*node, 0)})
	} else {
		fileSz, _ := strconv.Atoi(parts[0])
		n.children = append(n.children, &node{file, parts[1], fileSz, n, nil})
	}
}

func cd(wd *node, path string) *node {
	if path == ".." {
		return wd.parent
	}

	if path == "/" {
		cur := wd
		for cur.parent != nil {
			cur = cur.parent
		}
		return cur
	}

	for _, child := range wd.children {
		if child.name == path {
			return child
		}
	}

	return nil
}

func getSizes(n *node) (size int, subdirSizes []int) {
	subdirSizes = make([]int, 0)

	for _, child := range n.children {
		if child.kind == file {
			size += child.size
		} else {
			dirSize, subSubdirSizes := getSizes(child)
			subdirSizes = append(subdirSizes, subSubdirSizes...)
			subdirSizes = append(subdirSizes, dirSize)
			size += dirSize
		}
	}

	return size, subdirSizes
}

func P1(input []string) int {
	wd := &node{dir, "/", 0, nil, make([]*node, 0)}

	for _, input := range input {
		parts := strings.Split(input, " ")

		switch parts[1] {
		case "ls":
			continue
		case "cd":
			wd = cd(wd, parts[2])
		default:
			wd.addChildNode(parts)
		}
	}

	wd = cd(wd, "/")
	size, subdirSizes := getSizes(wd)
	dirSizes := append(subdirSizes, size)

	totalSize := 0
	for _, size := range dirSizes {
		if size <= 100000 {
			totalSize += size
		}
	}

	// fmt.Printf("size: %d\n", size)
	// fmt.Printf("subdirs: %v\n", subdirSizes)

	return totalSize
}

func P2(input []string) int {
	return 0
}
