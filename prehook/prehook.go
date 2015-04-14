package prehook

import (
	"fmt"

	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["prehook"] = Run
}

func Run(config *entrykit.Config) {
	//defer entrykit.Exec(config.Exec)
	entrykit.Error(fmt.Errorf("Not implemented"))
}
