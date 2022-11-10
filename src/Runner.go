package src

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"

	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/widget"
	//"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/widget"
)

func Runner(entry *widget.Entry, startcommand string) (io.ReadCloser, *bufio.Writer) {
	cmd := exec.Command(startcommand)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	readerr := bufio.NewReader(stdout)
	writer := bufio.NewWriter(stdin)
	go func() {
		GetOutput(readerr, entry)
	}()
	//writer.Write()
	return stdout, writer
}

func GetOutput(reader *bufio.Reader, entry *widget.Entry) string {
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

		//fmt.Print(output) //输出屏幕内容
		sumOutput += output
		entry.SetText(sumOutput)
		//entry.Refresh()
	}
	return sumOutput
}

func JavaParser() {

}
