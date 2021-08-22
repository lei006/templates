package cmd_unit

import (
	"io"
	"os/exec"
	"regexp"
	"strings"
)

type CmdUnit struct {
	std_in  io.WriteCloser
	Running bool

	cmd *exec.Cmd
}

func (this *CmdUnit) Start(cmd_str string) error {

	reg := regexp.MustCompile(`/\s*(".+?"|[^:\s])+((\s*:\s*(".+?"|[^\s])+)|)|(".+?"|[^"\s])+`)
	params := reg.FindAllString(cmd_str, -1)

	cmd := exec.CommandEx(params[0], strings.Join(params[1:], " "))
	tmp_std_in, err := cmd.StdinPipe()
	if err != nil {
		this.Running = false
		return err
	}

	err = cmd.Start()
	if err != nil {
		this.Running = false
		return err
	}

	this.std_in = tmp_std_in
	this.Running = true
	this.cmd = cmd

	return nil
}

func (this *CmdUnit) Wait() error {
	return this.cmd.Wait()
}

func (this *CmdUnit) Stop() {
	this.std_in.Write([]byte("q"))
}
