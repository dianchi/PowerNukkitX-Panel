package src

import (
	//"bufio"
	//"fmt"
	//"io"
	//"log"
	"os/exec"
	//"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/widget"
	//"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/widget"
)

func Runner(startcommand string) {
	cmd := exec.Command("cmd.exe", "/c", "start "+startcommand)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	/*
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
			GetOutput(readerr)
		}()
		//writer.Write()
		return stdout, writer
	*/
}

/*
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

		//fmt.Print(output) //输出屏幕内容
		sumOutput += output
	}
	return sumOutput
}

*/
