package segment

import (
	"log"
	"os"
	"os/user"
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"testing"
)


func Test_ReadOnlyStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := ReadOnlySegment("/usr", 0, conf, &buffer)
	output := buffer.String()
	expected := `\[\e[48;5;1m\]\[\e[38;5;0m\]⮀\[\e[38;5;15m\]\[\e[48;5;1m\] RO `
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != 1 {
		t.Errorf("Got: %d, Expected: %d", bg, 1)
	}
}


func Test_UserWriteOnlyStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	bg := ReadOnlySegment(usr.HomeDir, 0, conf, &buffer)
	output := buffer.String()
	expected := ""
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != 0 {
		t.Errorf("Got: %d, Expected: %d", bg, 0)
	}
}


func Test_GroupWriteOnlyStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	tmp_dir := "/tmp/go_prompt_testing"
	log.Println(tmp_dir)
	err := os.Mkdir(tmp_dir, 0774)
	if err != nil {
		log.Println(err)
	}
	os.Chown(tmp_dir, os.Getuid(), os.Getgid())
	err = os.Chmod(tmp_dir, 0774)
	if err != nil {
		log.Fatalln(err)
	}
	defer os.Remove(tmp_dir)
	bg := ReadOnlySegment(tmp_dir, 0, conf, &buffer)
	output := buffer.String()
	expected := ""
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != 0 {
		t.Errorf("Got: %d, Expected: %d", bg, 0)
	}

}

func Test_WriteableStringBash(t *testing.T) {
	var buffer bytes.Buffer
	conf := config.ConfigFromDefaults()
	bg := ReadOnlySegment("/tmp", 0, conf, &buffer)
	output := buffer.String()
	expected := ""
	if output != expected {
		t.Errorf("Got: \"%s\", Expected: \"%s\"", output, expected)
	}

	if bg != 0 {
		t.Errorf("Got: %d, Expected: %d", bg, 0)
	}
}