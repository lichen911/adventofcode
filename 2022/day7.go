package main

import (
	"bufio"
	"fmt"
	"os"
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

func (dir *Dir) calcSize() {
	for _, subDir := range dir.Dirs {
		dir.Size += subDir.Size
	}
	for _, file := range dir.Files {
		dir.Size += file.Size
	}
}

func main() {
	file, err := os.Open("day7_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := Dir{Name: "root"}
	currentDir := &root
	listDir := false
	var auditDirs []*Dir

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			// process user command
			listDir = false
			cmd := strings.Split(line, " ")[1:]
			switch cmd[0] {
			case "ls":
				listDir = true
			case "cd":
				cdArg := cmd[1]
				if cdArg == ".." {
					currentDir.calcSize()
					if currentDir.Size <= 100000 {
						auditDirs = append(auditDirs, currentDir)
					}

					currentDir = currentDir.ParentDir
				} else {
					for _, dir := range currentDir.Dirs {
						if dir.Name == cdArg {
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
			} else {
				fileSize, err := strconv.Atoi(strings.Split(line, " ")[0])
				checkError(err)
				fileName := strings.Split(line, " ")[1]
				currentDir.Files = append(currentDir.Files, File{Size: fileSize, Name: fileName})
			}
		} else {
			panic("Unexpected input")
		}
	}

	total := 0
	for _, dir := range auditDirs {
		total += dir.Size
	}
	fmt.Println("Total size:", total)
}
