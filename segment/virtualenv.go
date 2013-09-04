package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/prompt"
	"os"
	"path/filepath"
)

func VirtualEnvSegment(last_bg int, config prompt.Config, buffer *bytes.Buffer) int {
	env := os.Getenv("VIRTUAL_ENV")
	if env != "" {
		virt := filepath.Base(env)
		prompt.SegmentConnector(last_bg, config.Theme["virtual_env_bg"], config, buffer)
		buffer.WriteString(prompt.Color(38, config.Theme["virtual_env_fg"]))
		buffer.WriteString(prompt.Color(48, config.Theme["virtual_env_bg"]))
		buffer.WriteString(fmt.Sprintf(" %s ", virt))
		return config.Theme["virtual_env_bg"]
	}
	return last_bg
}
