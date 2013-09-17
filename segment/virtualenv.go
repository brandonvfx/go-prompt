package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
	"os"
	"path/filepath"
)

func VirtualEnvSegment(last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	env := os.Getenv("VIRTUAL_ENV")
	if env != "" {
		virt := filepath.Base(env)
		prompt.SegmentConnector(last_bg, conf.Theme["virtual_env_bg"], conf, buffer)
		buffer.WriteString(prompt.Color(38, conf.Theme["virtual_env_fg"]))
		buffer.WriteString(prompt.Color(48, conf.Theme["virtual_env_bg"]))
		buffer.WriteString(fmt.Sprintf(" %s ", virt))
		return conf.Theme["virtual_env_bg"]
	}
	return last_bg
}
