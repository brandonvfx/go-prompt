package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/config"
	"github.com/brandonvfx/go-prompt/prompt"
	"log"
	"os/exec"
	"strings"
)

func RvmSegment(last_bg int, conf config.Config, buffer *bytes.Buffer) int {
	rvm_cmd := exec.Command("rvm-prompt", "g")
	var rvm_out bytes.Buffer
	rvm_cmd.Stdout = &rvm_out
	err := rvm_cmd.Run()
	if err != nil {
		log.Println(err)
		return last_bg
	}

	rvm_env := strings.TrimSpace(rvm_out.String())
	if rvm_env != "" {
		rvm_env = strings.Replace(rvm_env, "@", "", -1)
		prompt.SegmentConnector(last_bg, conf.Theme["rvm_bg"], conf, buffer)

		buffer.WriteString(prompt.Color(38, conf.Theme["rvm_fg"]))
		buffer.WriteString(prompt.Color(48, conf.Theme["rvm_bg"]))
		buffer.WriteString(fmt.Sprintf(" %s ", rvm_env))

		return conf.Theme["rvm_bg"]
	}
	return last_bg
}
