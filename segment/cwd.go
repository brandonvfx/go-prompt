package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/prompt"
	"strings"
)

func CwdSegment(cwd string, last_bg int, config prompt.Config, buffer *bytes.Buffer) int {
	path_parts := strings.Split(cwd, "/")
	path_len := len(path_parts)

	maxdepth := config.Maxdepth
	var path []string
	if path_len > maxdepth {
		end_idx := path_len + (2 - maxdepth)
		path = append(path, path_parts[:2]...)
		path = append(path, config.Symbols["compressed_path"])
		path = append(path, path_parts[end_idx:]...)
		path_len = len(path)
	} else {
		path = path_parts[:]
	}
	
	prompt.SegmentConnector(last_bg, config.Theme["path_bg"], config, buffer)
	for i, path_part := range path {
		if i+1 != path_len {
			buffer.WriteString(prompt.Color(38, config.Theme["path_fg"]))
			buffer.WriteString(prompt.Color(48, config.Theme["path_bg"]))
			buffer.WriteString(fmt.Sprintf(" %s ", path_part))
			buffer.WriteString(prompt.Color(48, config.Theme["path_bg"]))
			buffer.WriteString(prompt.Color(38, config.Theme["separator_fg"]))
			buffer.WriteString(fmt.Sprintf("%s", config.Symbols["separator_thin"]))
		} else {
			buffer.WriteString(prompt.Color(38, config.Theme["cwd_fg"]))
			buffer.WriteString(prompt.Color(48, config.Theme["path_bg"]))
			buffer.WriteString(fmt.Sprintf(" %s ", path_part))
		}
	}
	return config.Theme["path_bg"]
}
