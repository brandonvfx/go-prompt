package segment

import (
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"os/user"
	"testing"
)

func Test_CwdHomeDirShortStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	usr, err := user.Current()
	if err != nil {
		t.Errorf(err.Error())
	}
	var cwd string

	cwd += usr.HomeDir
	cwd += "/go_prompt_testing"
	bg := CwdSegment(cwd, 0, conf, &buffer)
	output := buffer.String()

	expected := `\[\e[48;5;237m\]\[\e[38;5;0m\]⮀\[\e[38;5;250m\]\[\e[48;5;237m\] ~ \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;254m\]\[\e[48;5;237m\] go_prompt_testing `

	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["path_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["path_bg"])
	}
}

func Test_CwdHomeDirLongStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	usr, err := user.Current()
	if err != nil {
		t.Errorf(err.Error())
	}
	var cwd string

	cwd += usr.HomeDir
	cwd += "/go/prompt/testing/should/be/long/enough"
	bg := CwdSegment(cwd, 0, conf, &buffer)
	output := buffer.String()

	expected := `\[\e[48;5;237m\]\[\e[38;5;0m\]⮀\[\e[38;5;250m\]\[\e[48;5;237m\] ~ \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] go \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] … \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] be \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] long \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;254m\]\[\e[48;5;237m\] enough `

	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["path_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["path_bg"])
	}
}

func Test_CwdNonHomeDirShortStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()

	cwd := "/usr/local/go_prompt_testing"
	bg := CwdSegment(cwd, 0, conf, &buffer)
	output := buffer.String()

	expected := `\[\e[48;5;237m\]\[\e[38;5;0m\]⮀\[\e[38;5;250m\]\[\e[48;5;237m\] usr \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] local \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;254m\]\[\e[48;5;237m\] go_prompt_testing `

	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["path_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["path_bg"])
	}
}

func Test_CwdNonHomeDirLongStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()

	cwd := "/usr/local/go/prompt/testing/should/be/long/enough"
	bg := CwdSegment(cwd, 0, conf, &buffer)
	output := buffer.String()

	expected := `\[\e[48;5;237m\]\[\e[38;5;0m\]⮀\[\e[38;5;250m\]\[\e[48;5;237m\] usr \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] local \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] … \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] be \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;250m\]\[\e[48;5;237m\] long \[\e[48;5;237m\]\[\e[38;5;244m\]⮁\[\e[38;5;254m\]\[\e[48;5;237m\] enough `

	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != conf.Theme["path_bg"] {
		t.Errorf("Got: %d, Expected: %d", bg, conf.Theme["path_bg"])
	}
}

func Benchmark_CwdHomeDirShortString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	usr, err := user.Current()
	if err != nil {
		b.Errorf(err.Error())
	}
	var cwd string

	cwd += usr.HomeDir
	cwd += "/go_prompt_testing"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CwdSegment(cwd, 0, conf, &buffer)
	}
}

func Benchmark_CwdHomeDirLongString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	usr, err := user.Current()
	if err != nil {
		b.Errorf(err.Error())
	}
	var cwd string

	cwd += usr.HomeDir
	cwd += "/go/prompt/testing/should/be/long/enough"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CwdSegment(cwd, 0, conf, &buffer)
	}
}

func Benchmark_CwdNonHomeDirShortString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()

	cwd := "/usr/local/go_prompt_testing"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CwdSegment(cwd, 0, conf, &buffer)
	}
}

func Benchmark_CwdNonHomeDirLongString(b *testing.B) {
	b.StopTimer()
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()

	cwd := "/usr/local/go/prompt/testing/should/be/long/enough"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CwdSegment(cwd, 0, conf, &buffer)
	}
}
