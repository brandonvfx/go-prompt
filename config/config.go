package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const ConfigFile = "$HOME/.go_prompt"

var defaultConfig = `{
    "maxdepth": 5,
    "symbols": {
        "separator": "\u2B80",
        "separator_thin": "\u2B81",
        "compressed_path": "\u2026",
        "lock": "RO"
    },
    "theme": {
        "path_bg": 237,
        "path_fg": 250,
        "cwd_fg": 254, 
        "separator_fg": 244,
        "cmd_passed_bg": 236, 
        "cmd_passed_fg": 15,
        "cmd_failed_bg": 161,
        "cmd_failed_fg": 15,
        "cmd_fg": 15,
	    "virtual_env_bg": 35, 
        "virtual_env_fg": 0,
	    "repo_clean_bg": 148, 
        "repo_clean_fg": 0,
        "repo_dirty_bg": 161, 
        "repo_dirty_fg": 15,
        "rvm_bg": 44, 
        "rvm_fg": 0,
        "read_only_bg": 1,
        "read_only_fg": 15
    },
    "settings": {
        "CwdSegment": {
            "maxdepth": 5,
            "cwd_only": true
        }
    },
    "segments": [ "RvmSegment", "VirtualEnvSegment", "CwdSegment", "GitSegment", "CmdSegment" ]
}`

type Config struct {
	Symbols  map[string]string                 `json:symbols`
	Theme    map[string]int                    `json:theme`
	Settings map[string]map[string]interface{} `json:settings`
	Maxdepth int                               `json:maxdepth`
	Segments []string                          `json:segments`
}

func LoadConfig() Config {
	var file []byte
	var err error

	expandedConfigFile := os.ExpandEnv(ConfigFile)
	file, err = ioutil.ReadFile(expandedConfigFile)

	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	if err != nil && os.IsNotExist(err) {
		file = []byte(defaultConfig)
	}

	var conf Config
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalln(err)
	}

	return conf
}

func WriteDefaultConfig() error {
	expandedConfigFile := os.ExpandEnv(ConfigFile)
	_, err := ioutil.ReadFile(expandedConfigFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err == nil {
		return fmt.Errorf("Config file already exists: %s", expandedConfigFile)
	}

	err = ioutil.WriteFile(expandedConfigFile, []byte(defaultConfig), 0644)
	if err != nil {
		return err
	}
	log.Printf("Writing Config: %s", expandedConfigFile)
	return nil

}
