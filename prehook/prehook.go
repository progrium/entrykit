package prehook

import (
	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["prehook"] = Run
}

func Run(config *entrykit.Config) {
	defer entrykit.Exec(config.Exec)
	for _, task := range config.Tasks {
		cmd := entrykit.CommandTask(task)
		err := cmd.Run()
		if err != nil {
			entrykit.Error(err)
		}
	}
}
