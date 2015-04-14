package entrykit

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mgood/go-posix"
)

var Version string

var Cmds = make(map[string]func(config *Config))

var runlist = []string{
	//"prehook",
	"render",
	"switch",
	//"posthook",
	"codep",
	//"waitgrp",
}

func init() {
	Cmds[""] = Run
}

func Run(config *Config) {
	if len(config.Tasks) > 0 {
		Error(fmt.Errorf("Entrykit cannot take tasks via arguments"))
	}
	defer Exec(config.Exec)
	for _, name := range runlist {
		cmd, exists := Cmds[name]
		if exists {
			cmd(NewConfig(name, false))
		}
	}
}

func ProxySignals(tasks []*exec.Cmd) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)
	go func() {
		for sig := range signals {
			for _, t := range tasks {
				t.Process.Signal(sig)
			}
		}
	}()
}

func Exec(args []string) {
	if len(args) == 0 {
		return
	}
	bin, err := exec.LookPath(args[0])
	if err != nil {
		Error(err)
	}
	// todo: posthook non-exec mode
	for i := range args {
		arg, err := posix.ExpandEnv(args[i])
		if err == nil {
			args[i] = arg
		}
	}
	err = syscall.Exec(bin, args, os.Environ())
	if err != nil {
		Error(err)
	}
}

func ExecTask(task string) {
	args := strings.Fields(task)
	Exec(args)
}

func CommandTask(task string) *exec.Cmd {
	cmdSplit := strings.Fields(task)
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func Symlink() {
	for name, _ := range Cmds {
		os.Symlink(os.Args[0], filepath.Dir(os.Args[0])+"/"+name)
	}
}

func RanAs() string {
	cmdRun := filepath.Base(os.Args[0])
	cmd := ""
	for name := range Cmds {
		if name == cmdRun {
			cmd = name
		}
	}
	return cmd
}

func Error(err error) {
	fmt.Fprintln(os.Stderr, "!!", err)
	os.Exit(1)
}
