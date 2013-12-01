package prompt

import (
	// "fmt"
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"testing"
)

func Test_ColorString(t *testing.T) {
	output := Color(48, 35)
	expected := `\[\e[48;5;35m\]`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}
}

func Test_SegmentConnectorBlackBgBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	SegmentConnector(0, conf.Theme["virtual_env_bg"], conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;35m\]\[\e[38;5;0m\]⮀`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}
}

func Test_SegmentConnectorNonBlackBgBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	SegmentConnector(conf.Theme["rvm_bg"], conf.Theme["virtual_env_bg"], conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;35m\]\[\e[38;5;44m\]⮀`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}
}

func Test_SegmentConnectorBelowZeroBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	SegmentConnector(-1, conf.Theme["virtual_env_bg"], conf, &buffer)
	output := buffer.String()
	expected := ""
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}
}
