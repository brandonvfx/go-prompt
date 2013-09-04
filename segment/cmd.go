package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/prompt"
)

func get_status_color(status int, config prompt.Config) int {
	if status != 0 {
		return config.Theme["cmd_failed_bg"]
	}
	return config.Theme["cmd_passed_bg"]
}

func CmdSegment(return_code int, last_bg int, config prompt.Config, buffer *bytes.Buffer) int{
	status_bg := get_status_color(return_code, config)

	prompt.SegmentConnector(last_bg, status_bg, config, buffer)
	buffer.WriteString(prompt.ColorReset)

	buffer.WriteString(prompt.Color(48, status_bg))
	buffer.WriteString(prompt.Color(38, config.Theme["cmd_fg"]))
	buffer.WriteString(" $ ")

	buffer.WriteString(prompt.ColorReset)
	buffer.WriteString(prompt.Color(38, status_bg))
	buffer.WriteString(fmt.Sprintf("%s", config.Symbols["separator"]))
	
	return status_bg
}
