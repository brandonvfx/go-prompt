package prompt

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
)

const ColorTmpl = "\\[\\e%s\\]"
const ColorReset = "\\[\\e[0m\\]"

func Color(prefix int, code int) string {
	return fmt.Sprintf(ColorTmpl, fmt.Sprintf("[%d;5;%dm", prefix, code))
}

func SegmentConnector(last_bg int, segment_bg int, conf config.Config, buffer *bytes.Buffer) {
	if last_bg >= 0 {
		buffer.WriteString(Color(48, segment_bg))
		buffer.WriteString(Color(38, last_bg))
		buffer.WriteString(fmt.Sprintf("%s", conf.Symbols["separator"]))
	}
}
