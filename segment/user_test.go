package segment

import (
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"testing"
)

func Test_UserStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := UserSegment(0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;35m\]\[\e[38;5;0m\]⮀\[\e[38;5;0m\]\[\e[48;5;35m\] \u `
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["virtual_env_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["virtual_env_bg"])
	}
}

func Benchmark_UserString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		UserSegment(0, conf, &buffer)
	}
}
