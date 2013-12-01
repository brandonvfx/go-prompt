package segment

import (
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"testing"
)

func Test_CmdFailBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := CmdSegment(1, 0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;161m\]\[\e[38;5;0m\]⮀\[\e[0m\]\[\e[48;5;161m\]\[\e[38;5;15m\] $ \[\e[0m\]\[\e[38;5;161m\]⮀`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["cmd_failed_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["cmd_failed_bg"])
	}
}

func Test_CmdFailBash_2(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := CmdSegment(2, 0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;161m\]\[\e[38;5;0m\]⮀\[\e[0m\]\[\e[48;5;161m\]\[\e[38;5;15m\] $ \[\e[0m\]\[\e[38;5;161m\]⮀`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["cmd_failed_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["cmd_failed_bg"])
	}
}

func Test_CmdSuccessfulBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := CmdSegment(0, 0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;236m\]\[\e[38;5;0m\]⮀\[\e[0m\]\[\e[48;5;236m\]\[\e[38;5;15m\] $ \[\e[0m\]\[\e[38;5;236m\]⮀`
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["cmd_passed_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["cmd_passed_bg"])
	}
}

func Benchmark_CmdSuccessfulString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CmdSegment(0, 0, conf, &buffer)
	}
}

func Benchmark_CmdFailString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CmdSegment(1, 0, conf, &buffer)
	}
}
