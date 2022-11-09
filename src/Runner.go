package src

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	//"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/widget"
)

func Runner() (io.ReadCloser, io.WriteCloser) {
	cmd := exec.Command("tree")

	stdout, err := cmd.StdoutPipe()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	readerr := bufio.NewReader(stdout)
	go func() {
		GetOutput(readerr)
	}()
	return stdout, stdin
}

func GetOutput(reader *bufio.Reader) string {
	var sumOutput string
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		fmt.Print(output) //输出屏幕内容
		sumOutput += output
	}
	return sumOutput
}

func JavaParser() {

}
