package zombie

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func readArchive(path string) (lines []string, err error) {
	arch, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer arch.Close()

	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func cleanCommand(line string) (commands []string, err error) {
	var quotedString string
	var startString int
	var quote byte
	comm := strings.Split(line, " ")
	for i, l := 0, len(comm); i < l; i++ {
		item := comm[i]

		if quotedString != "" {
			quotedString = fmt.Sprintf("%s %s", quotedString, item)
		}

		if matched, err := regexp.MatchString("^('|\")", item); err != nil {
			return nil, err
		} else if matched && quotedString == "" {
			startString = i
			quotedString = item
			quote = quotedString[0]
		} else if matched, err := regexp.MatchString("[^\\\\]('|\")$", item); err != nil {
			return nil, err
		} else if currentQuote := item[len(item)-1]; matched && currentQuote == quote {
			comm = append(comm[:startString], comm[i:]...)
			comm[startString] = quotedString
			quotedString = ""
			quote = 0
			l = len(comm)
		}
	}

	return comm, nil
}

// ExecCommandPath ...
func ExecCommandPath(path string) (outputs []string, err error) {
	commands, err := readArchive(path)
	if err != nil {
		return nil, err
	}
	for _, command := range commands {
		param, err := cleanCommand(command)
		out, err := exec.Command(param[0], param[1:]...).CombinedOutput()
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, string(out))
	}
	return outputs, nil
}
