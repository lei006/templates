package cmd_unit

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

func RunWait(cmd_str string) error {

	var cmdUnit CmdUnit
	err := cmdUnit.Start(cmd_str)
	if err != nil {
		return err
	}

	err = cmdUnit.Wait()

	return err
	/*
		reg := regexp.MustCompile(`/\s*(".+?"|[^:\s])+((\s*:\s*(".+?"|[^\s])+)|)|(".+?"|[^"\s])+`)
		params := reg.FindAllString(cmd_str, -1)

		cmd := exec.Command(params[0], strings.Join(params[1:], " "))

		err := cmd.Start()
		if err != nil {
			return err
		}

		err = cmd.Wait()

		return err
	*/
}

func RunWaitCmd(cmd_str string) error {

	//Command("cmd", "/C", "rd /s/q D:\\1\\2")

	cmd := exec.Command("cmd.exe", "/c", cmd_str)

	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}

	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd error:", err.Error())
	}
	fmt.Println("cmd ", cmd_str)

	return err
}
