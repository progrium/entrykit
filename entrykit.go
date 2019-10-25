package entrykit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"

	"github.com/kardianos/osext"
	"github.com/mgood/go-posix"
)

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
	Cmds["entrykit"] = RunMulti
}

func RunMulti(config *Config) {
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
	// Split the task by words or quoted phrases using a regex. `strings.Fields`
	// splits by spaces including those in quotes, which breaks the command to be
	// run's argument processing.
	re := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	parts := re.FindAllString(task, -1)

	// And after that, exec.Command can't handle quoted arguments, with results
	// being a silent failure. Strip those out too.
	var cmdSplit []string
	for i := range parts {
		part := parts[i]
		if len(part) >= 2 {
			if part[0] == '\'' && part[len(part)-1] == '\'' {
				part = part[1 : len(part)-1]
			} else if part[0] == '"' && part[len(part)-1] == '"' {
				part = part[1 : len(part)-1]
			}
		}
		cmdSplit = append(cmdSplit, part)
	}

	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func Symlink() {
	binaryPath, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}
	for name, _ := range Cmds {
		target := filepath.Dir(binaryPath) + "/" + name
		fmt.Println("Creating symlink", target, "...")
		os.Symlink(os.Args[0], target)
	}
}

func RunCmd() {
	cmdRun := filepath.Base(os.Args[0])
	cmd := "entrykit"
	for name := range Cmds {
		if name == cmdRun {
			cmd = name
		}
	}
	Cmds[cmd](NewConfig(cmd, true))
}

func Error(err error) {
	fmt.Fprintln(os.Stderr, "!!", err)
	os.Exit(1)
}
