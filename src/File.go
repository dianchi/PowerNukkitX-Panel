package src

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func JarChecker(src string) bool {
	zr, err := zip.OpenReader(src)
	var rbool bool
	if err != nil {
		fmt.Println(err)
	}
	defer zr.Close()
	for _, f := range zr.File {
		if f.Name == "cn/nukkit/api/PowerNukkitXOnly.class" {
			rbool = true
		}
	}
	return rbool
}
func ServerChecker(path string) string {
	files := JarFinder(path)
	var filename string
	for _, files := range files {
		if JarChecker(files) {
			filename = files
			break
		}
	}
	return filename
}
func Unziper(src string, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		filePath := path.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := file.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()
			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
