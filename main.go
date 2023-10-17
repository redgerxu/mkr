package main

import (
	"fmt"
	"os"
	"strings"
)

func isFile(list []string) (bool, []string) {
	newList := make([]string, 0)

	found := false
	for _, item := range list {
		if item == "-f" || item == "--file" {
			found = true
			continue
		}

		newList = append(newList, item)
	}

	return found, newList
}

func main() {
	arguments := os.Args[1:]

	makeFile, args := isFile(arguments)

	segments := strings.Split(args[0], "/")

	cwd, err := os.Getwd()

	if args[0][0] == '/' {
		cwd = "/"
	}

	if err != nil {
		panic(err)
	}

	if segments[0] == "." {
		segments = segments[1:]
	} else if segments[1] == ".." {
		cwd += "/../"
	}

	file := ""

	if makeFile {
		file = segments[len(segments)-1]
		segments = segments[:(len(segments) - 1)]
	}

	dir := cwd + "/" + strings.Join(segments, "/")

	err = os.MkdirAll(dir, 0755)

	if err != nil {
		panic(err)
	}

	fmt.Println("Created directory " + dir)

	if makeFile {
		fn := dir + "/" + file
		os.Create(fn)
		fmt.Println("Created file " + fn)
	}
}
