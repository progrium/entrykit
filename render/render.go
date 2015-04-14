package render

import (
	"io/ioutil"

	"github.com/mgood/go-posix"
	"github.com/progrium/entrykit"
)

func init() {
	entrykit.Cmds["render"] = Run
}

func Run(config *entrykit.Config) {
	defer entrykit.Exec(config.Exec)
	for _, target := range config.Tasks {
		template := target + ".tmpl"
		data, err := ioutil.ReadFile(template)
		if err != nil {
			entrykit.Error(err)
		}
		render, err := posix.ExpandEnv(string(data))
		if err != nil {
			entrykit.Error(err)
		}
		// todo: use same filemode as template
		err = ioutil.WriteFile(target, []byte(render), 0644)
		if err != nil {
			entrykit.Error(err)
		}
	}
}
