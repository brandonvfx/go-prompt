package segment

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"syscall"
	//"reflect"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
)

func ReadOnlySegment(cwd string, last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return last_bg
	}
	cwd_stat, err := os.Stat(cwd)
	dir_uid := cwd_stat.Sys().(*syscall.Stat_t).Uid
	dir_gid := cwd_stat.Sys().(*syscall.Stat_t).Gid
	cwd_mode := cwd_stat.Mode()
	if err != nil {
		log.Println(err)
		return last_bg
	}

	if cwd_mode&0002 != 0 {
		return last_bg
	} else if cwd_mode&0020 != 0 && dir_gid == uint32(os.Getgid()) {
		return last_bg
	} else if cwd_mode&0200 != 0 && dir_uid == uint32(os.Getuid()) {
		return last_bg
	}

	prompt.SegmentConnector(last_bg, conf.Theme["read_only_bg"], conf, buffer)
	buffer.WriteString(prompt.Color(38, conf.Theme["read_only_fg"]))
	buffer.WriteString(prompt.Color(48, conf.Theme["read_only_bg"]))
	buffer.WriteString(fmt.Sprintf(" %s ", conf.Symbols["lock"]))
	return conf.Theme["read_only_bg"]
}
