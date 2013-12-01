package segment

import (
    "fmt"
    "log"
    "bytes"
    "github.com/brandonvfx/go-prompt/config"
    "testing"
    "os"
    "io/ioutil"
)

func MkFakeRvmPrompt(gemset string) {
    file_contents := fmt.Sprintf("#!/bin/bash\necho \"%s\"", gemset)
    os.Mkdir("/tmp/go_prompt_testing", 0755)
    err := ioutil.WriteFile("/tmp/go_prompt_testing/rvm-prompt", []byte(file_contents), 0755)
    if err != nil {
        log.Fatalln(err)
    }
}


func MkBrokenRvmPrompt() {
    file_contents := fmt.Sprintf("#!/bin/bash\nexit 1")
    os.Mkdir("/tmp/go_prompt_testing", 0755)
    err := ioutil.WriteFile("/tmp/go_prompt_testing/rvm-prompt", []byte(file_contents), 0755)
    if err != nil {
        log.Fatalln(err)
    }
}

func RmFakeRvmPrompt() {
    err := os.RemoveAll("/tmp/go_prompt_testing")
    if err != nil {
        log.Fatalln(err)
    }
}


func Test_RvmStringBash(t *testing.T) {
    var buffer bytes.Buffer
    conf := config.ConfigFromDefaults()
    MkFakeRvmPrompt("@go-prompt")
    defer RmFakeRvmPrompt()
    os.Setenv("PATH", "/tmp/go_prompt_testing:"+os.Getenv("PATH"))
    bg := RvmSegment(0, conf, &buffer)
    output := buffer.String()
    expected := `\[\e[48;5;44m\]\[\e[38;5;0m\]⮀\[\e[38;5;0m\]\[\e[48;5;44m\] go-prompt `
    if output != expected {
        t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
    }

    if bg != conf.Theme["rvm_bg"] {
        t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["rvm_bg"])
    }
}

func Test_RvmNoGemsetStringBash(t *testing.T) {
    var buffer bytes.Buffer
    conf := config.ConfigFromDefaults()
    MkFakeRvmPrompt("")
    defer RmFakeRvmPrompt()
    os.Setenv("PATH", "/tmp/go_prompt_testing:"+os.Getenv("PATH"))
    bg := RvmSegment(0, conf, &buffer)
    output := buffer.String()
    expected := ""
    if output != expected {
        t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
    }

    if bg != 0 {
        t.Errorf("Got: %d, Expected: %d", bg, 0)
    }
}


func Test_MissingRvmStringBash(t *testing.T) {
    var buffer bytes.Buffer
    conf := config.ConfigFromDefaults()
    os.Setenv("PATH", "")
    bg := RvmSegment(0, conf, &buffer)
    output := buffer.String()
    expected := ""
    if output != expected {
        t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
    }

    if bg != 0 {
        t.Errorf("Got: %d, Expected: %d", bg, 0)
    }
}


func Test_BrokenRvmStringBash(t *testing.T) {
    var buffer bytes.Buffer
    conf := config.ConfigFromDefaults()
    MkBrokenRvmPrompt()
    defer RmFakeRvmPrompt()
    os.Setenv("PATH", "/tmp/go_prompt_testing:"+os.Getenv("PATH"))
    bg := RvmSegment(0, conf, &buffer)
    output := buffer.String()
    expected := ""
    if output != expected {
        t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
    }

    if bg != 0 {
        t.Errorf("Got: %d, Expected: %d", bg, 0)
    }
}


