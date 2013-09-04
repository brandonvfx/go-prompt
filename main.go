package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brandonvfx/go-prompt/prompt"
	"github.com/brandonvfx/go-prompt/segment"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {

	cmd_return_code, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	expanded_config_file := os.ExpandEnv(prompt.ConfigFile)
	file, err := os.Open(expanded_config_file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var config prompt.Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if strings.HasPrefix(cwd, usr.HomeDir) {
		cwd = strings.Replace(cwd, usr.HomeDir, "~", 1)
	} else {
		cwd = strings.Replace(cwd, "/", "", 1)
	}

	var buffer bytes.Buffer

	last_bg := -1 //config.Theme["path_bg"]
	for _, segment_name := range config.Segments {
		switch segment_name {
		case "HostSegment":
			last_bg = segment.HostSegment(last_bg, config, &buffer)
		case "UserSegment":
			last_bg = segment.UserSegment(last_bg, config, &buffer)
		case "UserHostSegment":
			last_bg = segment.UserHostSegment(last_bg, config, &buffer)
		case "CwdSegment":
			last_bg = segment.CwdSegment(cwd, last_bg, config, &buffer)
		case "CmdSegment":
			last_bg = segment.CmdSegment(cmd_return_code, last_bg, config, &buffer)
		case "VirtualEnvSegment":
			last_bg = segment.VirtualEnvSegment(last_bg, config, &buffer)
		case "RvmSegment":
			last_bg = segment.RvmSegment(last_bg, config, &buffer)
		case "GitSegment":
			last_bg = segment.GitSegment(last_bg, config, &buffer)
		}
	}

	buffer.WriteString(prompt.ColorReset)
	buffer.WriteString("\n")
	fmt.Print(buffer.String())
}
