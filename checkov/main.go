package checkov

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrorInvalid = errors.New("Wrong string type")
)

type Check struct {
	info map[string]string
}

func ToVim(scanner *bufio.Scanner) []string {
	//errorfmt := `%f:%l: %m`
	titleRegex := regexp.MustCompile(`^Check: (?P<name>.*): "(?P<desc>.*)"`)
	fileRegex := regexp.MustCompile(`^\s*File: (?P<file>.*):(?P<start>.*\d*)-(\d*)$`)
	guideRegex := regexp.MustCompile(`^\s*Guide: (?P<guide>.*)`)
	statusRegex := regexp.MustCompile(`^\s*(?P<status>PASSED|FAILED) for resource: .*`)
	var ret []string

	var check Check
	for scanner.Scan() {
		line := scanner.Text()

		_ = check.Parse(titleRegex, line)
		_ = check.Parse(fileRegex, line)
		_ = check.Parse(guideRegex, line)
		_ = check.Parse(statusRegex, line)

		guide, guideFound := check.info["guide"]
		name, _ := check.info["name"]
		file, _ := check.info["file"]
		if guideFound && check.info["status"] == "FAILED" {
			err := fmt.Sprintf("%s:%s: %s %s %s",
				strings.TrimLeft(file, "/"), // checkov reports file names with a leading slash that will mess with vim
				check.info["start"], name, check.info["desc"], guide,
			)
			ret = append(ret, err)
		}

		if guideFound {
			check.info = nil
		}
	}

	return ret
}

func (c *Check) Parse(re *regexp.Regexp, line string) error {
	if c.info == nil {
		c.info = make(map[string]string)
	}

	match := re.FindStringSubmatch(line)
	if len(match) == 0 {
		return ErrorInvalid
	}

	for i, name := range match {
		if i != 0 && name != "" {
			c.info[re.SubexpNames()[i]] = match[i]
		}
	}
	return nil
}
