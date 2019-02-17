package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	//fileName := "birthday_001.txt"
	//newName, err := match(fileName, 4)
	//if err != nil {
	//	fmt.Println("no match")
	//	os.Exit(1)
	//}
	//fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	count := 0
	type rename struct{
		filename string
		path string
	}
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mv %s => %s\n", origPath, newPath)
	}
}

func match(fileName string, total int) (string, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces) - 1]
	fileName = strings.Join(pieces[0: len(pieces)-1], ".")
	pieces = strings.Split(fileName, "_")
	name := strings.Join(pieces[0: len(pieces) - 1], "_")
	number, err := strconv.Atoi(pieces[len(pieces) - 1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
	return "", nil
}
