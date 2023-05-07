package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type File struct {
	Size int
	Name string
}

type Dir struct {
	Name      string
	Files     []File
	Dirs      []*Dir
	ParentDir *Dir
	Size      int
}

func (dir *Dir) calcSizeRecurse() {
	for _, subDir := range dir.Dirs {
		subDir.calcSizeRecurse()
		dir.Size += subDir.Size
	}
	for _, file := range dir.Files {
		dir.Size += file.Size
	}
}

func (dir *Dir) findDirsLessThan(size int) []*Dir {
	var dirs []*Dir
	if dir.Size <= size {
		dirs = append(dirs, dir)
	}
	for _, subDir := range dir.Dirs {
		dirs = append(dirs, subDir.findDirsLessThan(size)...)
	}
	return dirs
}

func (dir *Dir) getAllDirSizes() []int {
	var sizes []int
	sizes = append(sizes, dir.Size)
	for _, subDir := range dir.Dirs {
		sizes = append(sizes, subDir.getAllDirSizes()...)
	}
	return sizes
}

func sumDirs(dirs []*Dir) int {
	var total int
	for _, dir := range dirs {
		total += dir.Size
	}
	return total
}

func main() {
	file, err := os.Open("day7_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := Dir{Name: "root"}
	currentDir := &root
	listDir := false

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			// process user command
			listDir = false
			cmd := strings.Split(line, " ")[1:]
			switch cmd[0] {
			case "ls":
				fmt.Println("Listing dir", currentDir.Name)
				listDir = true
			case "cd":
				cdArg := cmd[1]
				if cdArg == ".." {
					fmt.Println("Going up, from", currentDir.Name, "to", currentDir.ParentDir.Name)
					currentDir = currentDir.ParentDir
				} else {
					for _, dir := range currentDir.Dirs {
						if dir.Name == cdArg {
							fmt.Println("Going down, from", currentDir.Name, "to", dir.Name)
							currentDir = dir
							break
						}
					}
				}
			}
		} else if listDir {
			// read ls command output
			if strings.HasPrefix(line, "dir") {
				dirName := strings.Split(line, " ")[1]
				newDir := Dir{Name: dirName, ParentDir: currentDir, Size: 0}
				currentDir.Dirs = append(currentDir.Dirs, &newDir)
				fmt.Println("- dir", dirName)
			} else {
				fileSize, err := strconv.Atoi(strings.Split(line, " ")[0])
				checkError(err)
				fileName := strings.Split(line, " ")[1]
				currentDir.Files = append(currentDir.Files, File{Size: fileSize, Name: fileName})
				fmt.Println("- file", fileName, "size", fileSize)
			}
		} else {
			panic("Unexpected input")
		}
	}

	root.calcSizeRecurse()

	// Print solution to part 1
	smallDirs := root.findDirsLessThan(100000)
	fmt.Println("Total size of small dirs:", sumDirs(smallDirs))

	// Calculate solution to part 2
	unusedSpace := 70000000 - root.Size
	requiredSpace := 30000000 - unusedSpace
	var deleteDirSize int
	var allDirSizes []int

	fmt.Println("Root size:", root.Size)
	fmt.Println("Unused space:", unusedSpace, "Required space:", requiredSpace)

	allDirSizes = root.getAllDirSizes()
	sort.Ints(allDirSizes)
	// fmt.Println(allDirSizes)
	fmt.Println("All dir count:", len(allDirSizes))

	for _, size := range allDirSizes {
		if size >= requiredSpace {
			deleteDirSize = size
			break
		}
	}
	fmt.Println("Delete dir size:", deleteDirSize)
}
