not a proper readme, just my design notes so far

```
entrykit

Entrypoint tools for elegant, programmable containers
 - Useful for Docker, rkt, LXC, etc
 - Good for image authors. Eliminates helper/start scripts or depending on shells.
 - Can be good for image users. Allows users to program/extend containers.

/bin/entrykit --symlink

-e allow environment as well
-E allow environment prefixed with EK_
-f <file> use config file
-p prefix output of tasks

codep [-eE] [[name=]task...] [-- exec]
waitgrp [-eE] [[name=]task...] [-- exec]
render [-eE] [[name=]path...] [-- exec]
switch [-eE] [[name=]exec...] [-- exec]
prehook [-eE] [[name=]hook...] [-- exec]
posthook [-eE] [[name=]task...] [-- exec]
	disable exec, allows parent process to exist entire time
undaemon?


/bin/entrykit -f <file> -- <exec>
	prehook
	render
	switch
	posthook
	codep
	waitgrp

specific
	- no environment, inline args
	- no general entrykit, specific tools
general
	- open to environment
	- default entrykit

don't support && alternative to --
	even though it works without shell,
	-- behavior is not always equivalent of &&

intentionally no looping or restarting tool
	* primary use case is against best practice. use higher level supervisor/restart-policy
	* edge use cases are minimal and aren't worth encourage bad practice

SWITCH_SHELL=/bin/sh
RENDER_CONFIG=/config/consul.json
CODEP_NGINX=nginx -g
PREHOOK_HTPASSWD=htpasswd -bc /etc/nginx/htpasswd $HTPASSWD

versioning
	semver.
	pre 1.0: only minor is used (0.3.0)
	get to 1.0 quickly.
	major: stable interface for commands
	minor: compatible additions
	patch: compatible fixes

prehook -- /bin/consul agent -config-dir=${CONFIG_DIR:-/config}

render

split
join
replace
...most of strings pkg

https://github.com/teepark/envrender/blob/master/main.go


```
