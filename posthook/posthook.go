package posthook

import (
	"fmt"

	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["posthook"] = Run
}

func Run(config *entrykit.Config) {
	entrykit.Error(fmt.Errorf("Not implemented"))
}
