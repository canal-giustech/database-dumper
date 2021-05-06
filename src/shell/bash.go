package shell

import (
	"fmt"
	"os/exec"
	strings "strings"
)

func Execute(command string, parameters ...string) (string, error) {
	cmd := exec.Command(command, parameters...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Println(err.Error())
		return "", err
	}
	return string(stdout), err
}

func Md5SumFile(filename string) (string, error) {
	result,err:=Execute("md5sum", filename)
	if err != nil {
		return "", err
	}
	return strings.Split(result, " ")[0], nil
}