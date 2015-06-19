package main

import (
	"fmt"
	"os"

	"github.com/progrium/entrykit"

	_ "github.com/progrium/entrykit/codep"
	_ "github.com/progrium/entrykit/posthook"
	_ "github.com/progrium/entrykit/prehook"
	_ "github.com/progrium/entrykit/render"
	_ "github.com/progrium/entrykit/switch"
	_ "github.com/progrium/entrykit/waitgrp"
)

var Version string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
