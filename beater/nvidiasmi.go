package beater

import (
	"os/exec"
	"strconv"
	"strings"
)

type NVIDIACommands interface {
	GetSMI() (cmd *exec.Cmd)
	getNumGpus() int
}

type NvidiaCommand struct {
	query string
	env   string
}

//GetSMI nvidia-smi command
func (c NvidiaCommand) GetSMI() (cmd *exec.Cmd) {
	if c.env == "local" {
		return exec.Command("localnvidiasmi")
	}
	return exec.Command("nvidia-smi", "--query-gpu="+c.query, "--format=csv")
}

func (c NvidiaCommand) getNumGpus() int {
	if c.env == "local" {
		return 4
	}
	cmd := "ls /dev | grep nvidia | grep -v nvidia-uvm | grep -v nvidiactl | wc -l"
	out, err := exec.Command("bash", "-c", cmd).Output()
	ret := 0
	if err == nil {
		ret, _ = strconv.Atoi(strings.TrimSpace(string(out)))
	}
	return ret
}
