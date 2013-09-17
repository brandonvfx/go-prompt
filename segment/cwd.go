package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
	"log"
	"os/user"
	"strings"
)

func CwdSegment(cwd string, last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	if strings.HasPrefix(cwd, usr.HomeDir) {
		cwd = strings.Replace(cwd, usr.HomeDir, "~", 1)
	} else {
		cwd = strings.Replace(cwd, "/", "", 1)
	}

	path_parts := strings.Split(cwd, "/")
	path_len := len(path_parts)

	maxdepth := conf.Maxdepth
	var path []string
	if path_len > maxdepth {
		end_idx := path_len + (2 - maxdepth)
		path = append(path, path_parts[:2]...)
		path = append(path, conf.Symbols["compressed_path"])
		path = append(path, path_parts[end_idx:]...)
		path_len = len(path)
	} else {
		path = path_parts[:]
	}

	prompt.SegmentConnector(last_bg, conf.Theme["path_bg"], conf, buffer)
	for i, path_part := range path {
		if i+1 != path_len {
			buffer.WriteString(prompt.Color(38, conf.Theme["path_fg"]))
			buffer.WriteString(prompt.Color(48, conf.Theme["path_bg"]))
			buffer.WriteString(fmt.Sprintf(" %s ", path_part))
			buffer.WriteString(prompt.Color(48, conf.Theme["path_bg"]))
			buffer.WriteString(prompt.Color(38, conf.Theme["separator_fg"]))
			buffer.WriteString(fmt.Sprintf("%s", conf.Symbols["separator_thin"]))
		} else {
			buffer.WriteString(prompt.Color(38, conf.Theme["cwd_fg"]))
			buffer.WriteString(prompt.Color(48, conf.Theme["path_bg"]))
			buffer.WriteString(fmt.Sprintf(" %s ", path_part))
		}
	}
	return conf.Theme["path_bg"]
}
