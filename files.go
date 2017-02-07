package main

import (
	"fmt"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
)

type Dir struct {
	Root      string
	Dirs      []string
	Files     []string
	Images    []string
	AbsDirs   []string
	AbsFiles  []string
	AbsImages []string
}

func NewDir(path string) *Dir {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !fi.IsDir() {
		return nil
	} else {
		dir := Dir{Root: path}
		files, _ := ioutil.ReadDir(path)
		for _, fi := range files {
			absPath := fp.Join(path, fi.Name())
			relPath := fi.Name()
			if fi.IsDir() {
				dir.Dirs = append(dir.Dirs, relPath)
				dir.AbsDirs = append(dir.AbsDirs, absPath)
			} else {
				dir.Files = append(dir.Files, relPath)
				dir.AbsFiles = append(dir.AbsFiles, absPath)
				switch strings.ToLower(fp.Ext(relPath)) {
				case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
					dir.Images = append(dir.Images, relPath)
					dir.AbsImages = append(dir.AbsImages, absPath)
				default:
				}
			}
		}
		return &dir
	}
}

func hasPhoto(path string) bool {
	dir := NewDir(path)
	if len(dir.Images) > 0 {
		return true
	} else {
		for _, subpath := range dir.AbsDirs {
			if hasPhoto(subpath) {
				return true
			}
		}
	}
	return false
}
