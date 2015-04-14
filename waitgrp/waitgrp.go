package waitgrp

import (
	"fmt"

	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["waitgrp"] = Run
}

func Run(config *entrykit.Config) {
	entrykit.Error(fmt.Errorf("Not implemented"))
}
