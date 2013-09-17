package segment

import (
	"bytes"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
)

func HostSegment(last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	prompt.SegmentConnector(last_bg, conf.Theme["virtual_env_bg"], conf, buffer)
	buffer.WriteString(prompt.Color(38, conf.Theme["virtual_env_fg"]))
	buffer.WriteString(prompt.Color(48, conf.Theme["virtual_env_bg"]))
	buffer.WriteString(" \\h ")
	return conf.Theme["virtual_env_bg"]
}
