package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fileSystemFile struct {
	name string
	size int
}

type fileSystemDirectory struct {
	name     string
	DirSize  int
	files    []fileSystemFile
	children []*fileSystemDirectory
	parent   *fileSystemDirectory
}

func main() {
	fs := readShellData("input.txt")

	fmt.Printf("Part 1 - Sum of all directories of size 100000 or less: %d\n", walkFS(fs))

	smallerThan := getSmallerThan(fs, getRequiredSpace(fs))
	sort.Ints(smallerThan)

	fmt.Printf("Part 2 - Size of the directory required to be deleted for the update: %d\n", smallerThan[0])
}

func readShellData(input string) fileSystemDirectory {

	fs := fileSystemDirectory{}
	var currentDirectory *fileSystemDirectory

	shellData, err := os.Open(input)
	defer shellData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return fs

	}

	shellReader := bufio.NewReader(shellData)
	readingDirContents := false

	for {
		shellInfo, _, err := shellReader.ReadLine()

		if err == io.EOF {
			break
		}

		shellLine := string(shellInfo)
		cmd := ""
		cmdInput := ""

		// look for shell commands
		if strings.Contains(shellLine, "$") {
			readingDirContents = false
			cmd = shellLine[2:4]
			cmdInput = shellLine[4:]
			cmdInput = strings.TrimSpace(cmdInput)
		}

		// Handled a cd cmd
		if cmd == "cd" {

			if strings.Contains(cmdInput, "/") {
				// Handle root
				fs.name = "/"
				fs.parent = nil
				children := []*fileSystemDirectory{}
				fs.children = children
				fs.parent = nil
				currentDirectory = &fs

			} else if strings.Contains(cmdInput, "..") {
				// Handle going back one directory

				if currentDirectory.parent != nil {
					currentDirectory = currentDirectory.parent
				}

			} else {
				// Handle change to a child directory. The child should already exist in the child array because
				// it got created and added when the input was parsed

				cmdInput = strings.TrimSpace(cmdInput)

				if currentDirectory.children != nil {
					for _, child := range currentDirectory.children {

						if child.name == cmdInput {
							currentDirectory = child
						}
					}

				} else {
					fmt.Println(currentDirectory.name + " has no Children. This probably should never happen")
					fmt.Println("Trying to do: " + cmd + " " + cmdInput)
					fmt.Println()
				}

			}

		}

		if readingDirContents {
			// Add a new child directory
			if string(shellLine[:3]) == "dir" {
				child := fileSystemDirectory{name: shellLine[strings.Index(shellLine, " ")+1:], files: []fileSystemFile{}, parent: currentDirectory}
				currentDirectory.children = append(currentDirectory.children, &child)

			} else {
				// Otherwise the item is a file

				fileSizeS := shellLine[:strings.Index(shellLine, " ")+1]
				fileName := shellLine[strings.Index(shellLine, " ")+1:]

				fileSize, err := strconv.Atoi(strings.TrimSpace(fileSizeS))

				if err != nil {
					fmt.Println("There was an error converting the file size line: " + string(shellLine))
					fileSize = 0
				}

				currentDirectory.DirSize += fileSize

				file := fileSystemFile{name: fileName, size: fileSize}
				currentDirectory.files = append(currentDirectory.files, file)

			}
		}

		if cmd == "ls" {
			readingDirContents = true
		}

	}

	return fs
}

func getTreeSize(fsDir *fileSystemDirectory) int {

	tot := fsDir.DirSize

	for _, child := range fsDir.children {

		treesize := getTreeSize(child)
		tot += treesize

	}

	return tot
}

func walkFS(fsDir fileSystemDirectory) int {

	res := 0
	ts := getTreeSize(&fsDir)

	if ts <= 100000 {

		res = ts

	}

	for _, childDir := range fsDir.children {
		res += walkFS(*childDir)

	}

	return res
}

func getRequiredSpace(fsDir fileSystemDirectory) int {

	totalSpace := 70000000
	usedSpace := getTreeSize(&fsDir)

	requiredSpace := 30000000 - (totalSpace - usedSpace)

	return requiredSpace

}

func getSmallerThan(fsDir fileSystemDirectory, spaceRequired int) []int {
	sizes := []int{}

	ts := getTreeSize(&fsDir)

	if ts >= spaceRequired {

		sizes = append(sizes, ts)

	}
	for _, child := range fsDir.children {
		sizes = append(sizes, getSmallerThan(*child, spaceRequired)...)
	}

	return sizes

}
