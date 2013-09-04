package prompt

import (
	"bytes"
	"fmt"
)

const ColorTmpl = "\\[\\e%s\\]"
const ColorReset = "\\[\\e[0m\\]"
const ConfigFile = "$HOME/.go_powerline"

type Config struct {
	Symbols  map[string]string                 `json:symbols`
	Theme    map[string]int                    `json:theme`
	Settings map[string]map[string]interface{} `json:settings`
	Maxdepth int                               `json:maxdepth`
	Segments []string                          `json:segments`
}

func Color(prefix int, code int) string {
	return fmt.Sprintf(ColorTmpl, fmt.Sprintf("[%d;5;%dm", prefix, code))
}

func SegmentConnector(last_bg int, segment_bg int, config Config, buffer *bytes.Buffer) {
	if last_bg >= 0 {
		buffer.WriteString(Color(48, segment_bg))
		buffer.WriteString(Color(38, last_bg))
		buffer.WriteString(fmt.Sprintf("%s", config.Symbols["separator"]))
	}
}
