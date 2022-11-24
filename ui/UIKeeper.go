package ui

import (
	"io/ioutil"
)

func FileReader(path string) string {
	data, _ := ioutil.ReadFile(path)
	return string(data)
}

func IndexHtml() string {
	html := FileReader("./html/index.html")
	return html
}
