package switch_

import (
	"os"

	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["switch"] = Run
}

func Run(config *entrykit.Config) {
	if len(os.Args) < 2 {
		entrykit.Exec(config.Exec)
		return
	}
	last := os.Args[len(os.Args)-1]
	for name, task := range config.Tasks {
		if name == last {
			entrykit.ExecTask(task)
			return
		}
	}
	entrykit.Exec(config.Exec)
}
