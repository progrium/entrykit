package main

import (
	"github.com/progrium/entrykit"
	_ "github.com/progrium/entrykit/codep"
)

var cmd = "codep"

func main() {
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
