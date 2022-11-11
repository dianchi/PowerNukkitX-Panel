package src

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func JavaVersionChecker(originjavapath []string) string {
	var result string
	var out []byte
	//regex := "^java version \".*\""
	//reg := regexp.MustCompile(regex)
	for i, op := range originjavapath {
		op1 := strings.Replace(op, "\r", "", -1)
		op2 := strings.Replace(op1, "\n", "", -1)
		cmd := exec.Command(op2, "-version")
		out, _ = cmd.CombinedOutput()
		//version := reg.FindAllString(string(out), -1)
		if VersionGreaterThanT(string(out), "17.0.0") {
			result = string(originjavapath[i])
		}
	}
	return strings.Replace(result, "\n", "", -1)
}

func JavaWhere() []string {
	regex := ".*\\n"
	reg := regexp.MustCompile(regex)
	cmd := exec.Command("where", "java")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return reg.FindAllString(string(out), -1)
}
