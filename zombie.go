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

// func cleanLine(line string) (commands []string, err error) {
// 	outs := strings.Trim(line, " ")
// 	outsArr := strings.Split(outs, "'")

// 	for _, out := range outs {

// 	}
// 	return out, nil
// 	// strings.Split(command, " ")
// }

// ExecCommandPath ...
func ExecCommandPath(path string) (outputs []string, err error) {
	commands, err := readArchive(path)
	if err != nil {
		return nil, err
	}
	for _, command := range commands {
		// o, err := cleanLine(command)
		// fmt.Println(o)
		comm := strings.Split(command, " ")

		var aspas string
		var startString int
		var endString int
		for i, item := range comm {
			if matched, err := regexp.MatchString("^'", item); err != nil {
				fmt.Println(err.Error())
			} else if matched && aspas == "" {
				startString = i
				aspas = item
			} else if matched, err := regexp.MatchString("'$", item); err != nil {
				fmt.Println(err.Error())
			} else if matched {
				aspas = fmt.Sprintf("%s %s", aspas, item)
				endString = i
				break
			} else {
				if aspas != "" {
					aspas = fmt.Sprintf("%s %s", aspas, item)
				}
			}
		}
		if aspas != "" {
			comm = append(comm[:startString], comm[endString:]...)
			comm[startString] = aspas
		}

		fmt.Println(comm)

		out, err := exec.Command(comm[0], comm[1:]...).CombinedOutput()
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, string(out))
		// fmt.Println(outputs)
	}
	return outputs, nil
}
