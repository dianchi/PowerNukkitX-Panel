package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func ServerCore() ([]string, []string, []string) {
	client := http.DefaultClient
	resp, err := client.Get(ServerCoreURL)
	var name, url, time []string
	if err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < int(gjson.GetBytes(content, "#").Int()); i++ {
		name = append(name, gjson.GetBytes(content, strconv.Itoa(i)+".name").String())
		url = append(url, gjson.GetBytes(content, strconv.Itoa(i)+".url").String())
		time = append(time, gjson.GetBytes(content, strconv.Itoa(i)+".lastModified").String())
	}
	return name, url, time
}

func JVMList() ([]string, []string, []string) {
	client := http.DefaultClient
	resp, err := client.Get(JVMURL)
	var name, url, size []string
	if err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < int(gjson.GetBytes(content, "#").Int()); i++ {
		name = append(name, gjson.GetBytes(content, strconv.Itoa(i)+".name").String())
		url = append(url, gjson.GetBytes(content, strconv.Itoa(i)+".download.windows-x86.url").String())
		size = append(size, gjson.GetBytes(content, strconv.Itoa(i)+".download.windows-x86.size").String())
	}
	return name, url, size
}
