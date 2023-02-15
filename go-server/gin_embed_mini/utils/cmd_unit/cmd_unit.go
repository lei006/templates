package cmd_unit

import (
	"io"
	"os/exec"
	"regexp"
	"server/utils/exec_fix"
	"strings"
	"time"
)

type CmdUnit struct {
	std_in  io.WriteCloser
	Running bool

	cmd        *exec.Cmd
	start_time time.Time

	Id       int64
	FilePath string

	dogTime int64 //喂狗时间
}

func (this *CmdUnit) Start(cmd_str string) error {

	reg := regexp.MustCompile(`/\s*(".+?"|[^:\s])+((\s*:\s*(".+?"|[^\s])+)|)|(".+?"|[^"\s])+`)
	params := reg.FindAllString(cmd_str, -1)

	cmd := exec_fix.CommandEx(params[0], strings.Join(params[1:], " "))

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
	this.start_time = time.Now()
	return nil
}

func (this *CmdUnit) State() string {

	if this.cmd == nil {
		return ""
	}

	return this.cmd.ProcessState.String()
}

func (this *CmdUnit) TimeLen() int64 {
	return time.Now().Unix() - this.start_time.Unix()
}

func (this *CmdUnit) Stop() {

	this.Running = false
	this.std_in.Write([]byte("q"))
}

func (this *CmdUnit) Wait() error {
	return this.cmd.Wait()
}

func (this *CmdUnit) UpdateDogTime() {
	this.dogTime = time.Now().Unix() //喂狗时间
}

func (this *CmdUnit) GeteDogTime() int64 {
	return time.Now().Unix() - this.dogTime //喂狗多久了..
}
