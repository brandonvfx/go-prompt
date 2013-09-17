package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
)

func get_status_color(status int, conf config.Config) int {
	if status != 0 {
		return conf.Theme["cmd_failed_bg"]
	}
	return conf.Theme["cmd_passed_bg"]
}

func CmdSegment(return_code int, last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	status_bg := get_status_color(return_code, conf)

	prompt.SegmentConnector(last_bg, status_bg, conf, buffer)
	buffer.WriteString(prompt.ColorReset)

	buffer.WriteString(prompt.Color(48, status_bg))
	buffer.WriteString(prompt.Color(38, conf.Theme["cmd_fg"]))
	buffer.WriteString(" $ ")

	buffer.WriteString(prompt.ColorReset)
	buffer.WriteString(prompt.Color(38, status_bg))
	buffer.WriteString(fmt.Sprintf("%s", conf.Symbols["separator"]))

	return status_bg
}
