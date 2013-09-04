package segment

import (
	"bytes"
	"github.com/brandonvfx/go-prompt/prompt"
)

func UserHostSegment(last_bg int, config prompt.Config, buffer *bytes.Buffer) int {
	prompt.SegmentConnector(last_bg, config.Theme["virtual_env_bg"], config, buffer)
	buffer.WriteString(prompt.Color(38, config.Theme["virtual_env_fg"]))
	buffer.WriteString(prompt.Color(48, config.Theme["virtual_env_bg"]))
	buffer.WriteString(" \\u@\\h ")
	return config.Theme["virtual_env_bg"]
}
