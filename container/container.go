package container

import (
	"os/exec"
	"syscall"
	"os"
	"../flag"
)

func CreateContainer(name string) {

	cmd := exec.Command("/bin/bash", "-c", "sh");
	cmd.Stdin = os.Stdin;
	cmd.Stdout = os.Stdout;
	cmd.Stderr = os.Stderr;
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:syscall.CLONE_NEWUTS,
	}
	flag.Check(syscall.Sethostname([]byte(name)));
	flag.Check(cmd.Run());
}