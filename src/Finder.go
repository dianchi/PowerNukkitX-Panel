package src

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	//"strings"
)

func JarFinder(path string) []string {
	files, _ := GetAllFile(path)
	var filename []string
	for _, f := range files {
		suffix := filepath.Ext(f)
		if suffix == ".jar" {
			filename = append(filename, f)
		}
	}
	return filename
}

func GetAllFile(pathname string) ([]string, error) {
	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
