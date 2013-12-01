package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const DefaultConfigPath = "$HOME/.go_prompt"

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
	"segments": [ "UserSegment", "VirtualEnvSegment", "CwdSegment", "GitSegment", "CmdSegment" ]
}`

type Config struct {
	Symbols  map[string]string                 `json:symbols`
	Theme    map[string]int                    `json:theme`
	Settings map[string]map[string]interface{} `json:settings`
	Maxdepth int                               `json:maxdepth`
	Segments []string                          `json:segments`
}

func LoadConfig() Config {
	conf, err := ConfigFromFile(DefaultConfigPath)

	if err != nil {
		conf = ConfigFromDefaults()
	}

	return conf
}

func ConfigFromFile(filepath string) (Config, error) {
	var conf Config
	expandedConfigFile := os.ExpandEnv(filepath)
	conf_data, err := ioutil.ReadFile(expandedConfigFile)

	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
		return conf, errors.New("Error loading config file.")
	}

	if err != nil && os.IsNotExist(err) {
		return conf, errors.New("Config file not found")
	}

	err = json.Unmarshal(conf_data, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf, nil
}

func ConfigFromDefaults() Config {
	var conf Config
	conf_data := []byte(defaultConfig)
	err := json.Unmarshal(conf_data, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

func WriteDefaultConfig(filepath string) error {
	expandedConfigFile := os.ExpandEnv(filepath)
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
