# Entrykit

Entrypoint tools for elegant containers.

Entrykit takes common tasks you might put in an entrypoint start script and lets you quickly set them up in your Dockerfile. It's sort of like an init process,
but we don't believe in heavyweight init systems inside containers. Instead, Entrykit takes care of practical tasks for building minimal containers.

## Getting Entrykit

In your Dockerfile download a release of Entrykit and extract it into your
PATH, such as `/bin` for simplicity. For best experience, run `entrykit --symlink` to set up the subcommands as regular commands.

See the `example` directory for a demo Nginx container using Entrykit.

## Using Entrykit

Once set up, you use Entrykit commands in your entrypoint. You can use just one or chain them together. They all have the same usage structure:

    <command> [[name=]task...] [-- exec]

Commands are documented below. Commands take one or more optionally named "tasks", which is often a shell command, but is sometimes a file to operate on. Then `--` is used to define the next operation. This is similar to `&&` except it's built in to all commands so you don't need to run these commands in a subshell.

## Commands

### `codep` - Codependent tasks

`codep` runs multiple processes in parallel, proxying signals, but unlike nearly every init system, it kills all processes if one process terminates. This allows the container to exit so Docker or another init system can cleanly restart it if appropriate.

This is ideal for runtime configuration rendering tools, such as conf.d and consul-template, or anything else that makes sense to run as a co-process in the container. Don't go overboard!

```
ENTRYPOINT ["codep", \
    "/bin/config-reloader", \
    "/usr/sbin/nginx" ]
```

You can run more tasks after your tasks exit using `--`, but this is not terribly common.

### `render` - Template rendering

`render` takes one or more paths to files that will be rendered using [Sigil](https://github.com/gliderlabs/sigil) templating. The template is loaded from a file with the same path but with `.tmpl` added extension. For example, if you want to render a template at `/etc/nginx.conf`, then you would copy a template file to `/etc/nginx.conf.tmpl` and use `render /etc/nginx.conf`.

This is particularly useful to use environment variables in configuration, which is our preferred way to configure containers at boot time. But it also comes with the rest of Sigil's configuration oriented templating functions.

```
COPY ./nginx.conf.tmpl /etc/nginx.conf.tmpl
ENTRYPOINT ["render", "/etc/nginx.conf", "--", "/usr/sbin/nginx"]
```

Since you usually have more to do after the `render` command, it's typical to chain with `--`. Anything after `--` is exec'd into.

### `switch` - Command switching

`switch` allows you to exec into alternative processes than your normal entrypoint based on the command provided when the container is run. We typically like containers that need no command and just do their thing immediately, but sometimes there are alternative modes of operation such as getting into the shell or displaying version or help information.

This is the first command to really take advantage of named tasks. The name of the task is the command string it will switch on, and the value is the full command it will run. For example, you can expose the shell when users run with the command `shell` with `switch shell=/bin/sh`. And as usual you can provide multiple tasks for more than one command.

```
ENTRYPOINT ["switch", "shell=/bin/sh", "version=nginx -v", "--", "/usr/sbin/nginx"]
```
If none of the commands are provided, it goes on to exec the next task after `--`.

### `prehook` - Run pre-commands on start

If there are other set up tasks to perform, you can add them with `prehook`. You can specify multiple tasks and they'll be run in order. If they fail, the chained tasks will not be run. This is particularly interesting when you use an undocumented flag that allows users to specify their own prehook commands. This is an example of how Entrykit can be used to make your containers more customizable by the user. But for now, it's just a way to run serial tasks before your final entrypoint command.

Here we display the Nginx version before starting Nginx:
```
ENTRYPOINT ["prehook", "nginx -V", "--", "/usr/sbin/nginx"]
```
## Chaining

All these commands can be used together. Here is an example of all of them being used together as demonstrated in the example directory:

```
ENTRYPOINT [ \
  "switch", \
    "shell=/bin/sh", \
    "version=nginx -v", "--", \
  "render", "/demo/nginx.conf", "--", \
  "prehook", "nginx -V", "--", \
  "codep", \
    "/bin/reloader 3", \
    "/usr/sbin/nginx -c /demo/nginx.conf" ]
```

## Other ways to define your entrypoint

Although not documented or properly tested, there are other ways you can set up these entrypoint commands. One way is with environment variables defined previously in your Dockerfile. It would look something like this:

```
ENV SWITCH_SHELL=/bin/sh
ENV RENDER_CONFIG=/etc/nginx.conf
ENV CODEP_NGINX=nginx -g
ENV CODEP_CONFD=confd
ENV PREHOOK_HTPASSWD=htpasswd -bc /etc/nginx/htpasswd $HTPASSWD

ENTRYPOINT ["entrykit -e"]
```
There is potentially another flag implemented to read config like that from a file. However, and this might be desired, this opens up the ability for users to mess with your entrypoint! But it's only possible if explicitly enabled.

## License

MIT 
