package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
	"github.com/brandonvfx/go-prompt/segment"
	"log"
	"os"
	"strconv"
)

const Version = "0.1.0"

var writeConfig bool
var printVersion bool

func init() {
	flag.BoolVar(&writeConfig, "write-config", false, "Write default config to ~/.go_prompt")
	flag.BoolVar(&printVersion, "version", false, "Print version and exit")
	flag.BoolVar(&printVersion, "v", false, "Print version and exit")
}


func usage() {
	fmt.Println("go-prompt [-write-config] [-verison|-v] [command_exit_code]")
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if writeConfig {
		err := config.WriteDefaultConfig()
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
	
	if printVersion{
		fmt.Printf("go-prompt v%v\n", Version)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		log.Fatalln("No args passed")
	}

	cmd_return_code, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	conf := config.LoadConfig()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	var buffer bytes.Buffer

	last_bg := -1
	for _, segment_name := range conf.Segments {
		switch segment_name {
		case "HostSegment":
			last_bg = segment.HostSegment(last_bg, conf, &buffer)
		case "UserSegment":
			last_bg = segment.UserSegment(last_bg, conf, &buffer)
		case "UserHostSegment":
			last_bg = segment.UserHostSegment(last_bg, conf, &buffer)
		case "CwdSegment":
			last_bg = segment.CwdSegment(cwd, last_bg, conf, &buffer)
		case "CmdSegment":
			last_bg = segment.CmdSegment(cmd_return_code, last_bg, conf, &buffer)
		case "VirtualEnvSegment":
			last_bg = segment.VirtualEnvSegment(last_bg, conf, &buffer)
		case "RvmSegment":
			last_bg = segment.RvmSegment(last_bg, conf, &buffer)
		case "GitSegment":
			last_bg = segment.GitSegment(last_bg, conf, &buffer)
		case "ReadOnlySegment":
			last_bg = segment.ReadOnlySegment(cwd, last_bg, conf, &buffer)
		}
	}

	buffer.WriteString(prompt.ColorReset)
	buffer.WriteString("\n")
	fmt.Print(buffer.String())
}
