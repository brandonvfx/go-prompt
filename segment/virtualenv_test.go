package segment

import (
	// "fmt"
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"os"
	"testing"
)

func Test_VirtualenvEmptyBash(t *testing.T) {
	os.Setenv("VIRTUAL_ENV", "")
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	last_bg := 0
	bg := VirtualEnvSegment(last_bg, conf, &buffer)
	output := buffer.String()
	expected := ""
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != last_bg {
		t.Errorf("Got: %d, Expected: %d", bg, last_bg)
	}
}

func Test_VirtualenvStringBash(t *testing.T) {
	os.Setenv("VIRTUAL_ENV", "go-prompt")
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := VirtualEnvSegment(0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;35m\]\[\e[38;5;0m\]⮀\[\e[38;5;0m\]\[\e[48;5;35m\] go-prompt `
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["virtual_env_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["virtual_env_bg"])
	}
}

func Benchmark_VirtualenvString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VirtualEnvSegment(0, conf, &buffer)
	}
}
