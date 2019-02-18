package main

import (
	"flag"
	"fmt"
	"sort"

	//"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()

	walkDir := "sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)
		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i+1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, newPath, err.Error())
				}
			}
		}

	}
}

type matchResult struct {
	base string
	index int
	ext string
}

func match(fileName string) (*matchResult, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces) - 1]
	fileName = strings.Join(pieces[0: len(pieces)-1], ".")
	pieces = strings.Split(fileName, "_")
	name := strings.Join(pieces[0: len(pieces) - 1], "_")
	number, err := strconv.Atoi(pieces[len(pieces) - 1])
	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern", fileName)
	}
	return &matchResult{
		strings.Title(name),
		number,
		ext,
	}, nil
	//return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}
