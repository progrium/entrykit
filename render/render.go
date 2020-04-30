package render

import (
	"io/ioutil"
	"path/filepath"
	"os"
	"syscall"
	"github.com/gliderlabs/sigil"
	"github.com/progrium/entrykit"

	_ "github.com/gliderlabs/sigil/builtin"
)

func init() {
	entrykit.Cmds["render"] = Run
}

func Run(config *entrykit.Config) {
	defer entrykit.Exec(config.Exec)
	for name, target := range config.Tasks {
		template :=  filepath.Join(config.TemplateDir, target + ".tmpl")
		data, err := ioutil.ReadFile(template)
		if err != nil {
			entrykit.Error(err)
		}
		render, err := sigil.Execute(data, nil, name)
		if err != nil {
			entrykit.Error(err)
		}
		fi, err := os.Stat(template)
		if err != nil {
		  entrykit.Error(err)
		}
		var st syscall.Stat_t
		err = syscall.Stat(template, &st)
		if err != nil {
		  entrykit.Error(err)
		}
		err = ioutil.WriteFile(target, render.Bytes(), fi.Mode())
		if err != nil {
			entrykit.Error(err)
		}
		err = os.Chown(target, int(st.Uid), int(st.Gid))
		if err != nil {
			entrykit.Error(err)
		}
	}
}
