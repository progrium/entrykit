package entrykit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const namespacePrefix = "ek_"

type Config struct {
	Cmd    string
	Prefix bool
	Tasks  map[string]string
	Exec   []string
}

func NewConfig(cmd string, exec bool) *Config {
	config := &Config{
		Cmd:   cmd,
		Tasks: make(map[string]string),
	}
	var configFile string
	var useEnv, namespaceEnv bool
	var argTasks []string
	args := os.Args[1:]
Loop:
	for i, arg := range args {
		switch arg {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		case "-e":
			useEnv = true
		case "-E":
			useEnv = true
			namespaceEnv = true
		case "-p":
			config.Prefix = true
		case "-f":
			configFile = args[i+1]
		case configFile:
			continue
		case "--":
			if exec {
				config.Exec = args[i+1:]
			}
			break Loop
		default:
			if strings.HasPrefix(arg, "-") {
				Error(fmt.Errorf("Unknown option %s", arg))
			}
			argTasks = append(argTasks, arg)
		}
	}
	if configFile != "" {
		file, err := os.Open(configFile)
		if err != nil {
			Error(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			config.addTask(envTask(scanner.Text(), cmd, false))
		}
		if err := scanner.Err(); err != nil {
			Error(err)
		}
	}
	if useEnv || len(argTasks) == 0 {
		for _, line := range os.Environ() {
			config.addTask(envTask(line, cmd, namespaceEnv))
		}
	}
	for _, kvp := range argTasks {
		config.addTask(argTask(kvp))
	}
	return config
}

func (c *Config) addTask(kvp []string) {
	if len(kvp) == 2 {
		c.Tasks[kvp[0]] = kvp[1]
	}
}

func envTask(kvp, cmd string, namespace bool) []string {
	prefix := cmd + "_"
	if namespace {
		prefix = namespacePrefix + prefix
	}
	parts := strings.SplitN(kvp, "=", 2)
	if len(parts) < 2 || cmd == "" {
		return nil
	}
	key := strings.ToLower(parts[0])
	if !strings.HasPrefix(key, prefix) {
		return nil
	}
	key = strings.Replace(key, prefix, "", 1)
	return []string{key, parts[1]}
}

func argTask(kvp string) []string {
	parts := strings.SplitN(kvp, "=", 2)
	if len(parts) < 2 {
		// use base name of first arg without extension
		args := strings.SplitN(parts[0], " ", -1)
		filename := strings.SplitN(args[0], ".", 2)
		return []string{filepath.Base(filename[0]), parts[0]}
	} else {
		return []string{parts[0], parts[1]}
	}
}
