package main

import (
	"os"

	"github.com/progrium/entrykit"

	_ "github.com/progrium/entrykit/codep"
	_ "github.com/progrium/entrykit/posthook"
	_ "github.com/progrium/entrykit/prehook"
	_ "github.com/progrium/entrykit/render"
	_ "github.com/progrium/entrykit/switch"
	_ "github.com/progrium/entrykit/waitgrp"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "symlink" {
		entrykit.Symlink()
		return
	}
	cmd := entrykit.RanAs()
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
