package segment

import (
	"bytes"
	"fmt"
	"github.com/brandonvfx/go-prompt/prompt"
	"os/exec"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`Your branch is (?P<direction>ahead|behind).*?(?P<count>\d+) commit`)

func git_status() (clean bool, extra string) {
	git_cmd := exec.Command("git", "status", "--ignore-submodules")
	var git_out bytes.Buffer
	git_cmd.Stdout = &git_out
	err := git_cmd.Run()
	if err != nil {
		return true, ""
	}

	clean = true
	extra = ""

	if !strings.Contains(git_out.String(), "nothing to commit") {
		clean = false
	}

	if strings.Contains(git_out.String(), "Untracked files") {
		extra = " +"
	}

	matches := re.FindStringSubmatch(git_out.String())
	if matches != nil {
		var dir string
		if matches[1] == "ahead" {
			dir = "\u21E1"
		} else {
			dir = "\u21E3"
		}
		extra = fmt.Sprintf(" %s%s", matches[2], dir)
	}

	return clean, extra
}

func GitSegment(last_bg int, config prompt.Config, buffer *bytes.Buffer) int {
	git_cmd := exec.Command("git", "branch")
	var git_out bytes.Buffer
	git_cmd.Stdout = &git_out
	err := git_cmd.Run()
	if err != nil {
		return last_bg
	}

	grep_cmd := exec.Command("grep", "*")
	grep_cmd.Stdin = strings.NewReader(git_out.String())

	var grep_out bytes.Buffer
	grep_cmd.Stdout = &grep_out
	err = grep_cmd.Run()
	if err != nil {
		return last_bg
	}
	branch := grep_out.String()
	branch = strings.TrimSpace(strings.Replace(branch, "*", "", 1))

	clean, extra := git_status()

	var fg_color, bg_color int
	if clean {
		fg_color = config.Theme["repo_clean_fg"]
		bg_color = config.Theme["repo_clean_bg"]
	} else {
		fg_color = config.Theme["repo_dirty_fg"]
		bg_color = config.Theme["repo_dirty_bg"]
	}

	prompt.SegmentConnector(last_bg, bg_color, config, buffer)
	
	buffer.WriteString(prompt.Color(48, bg_color))
	buffer.WriteString(prompt.Color(38, fg_color))
	buffer.WriteString(fmt.Sprintf(" %s%s ", branch, extra))
	return bg_color
}
