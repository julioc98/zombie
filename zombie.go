package zombie

import (
	"bufio"
	"os"
	"os/exec"
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

// ExecCommandPath ...
func ExecCommandPath(path string) (outputs []string, err error) {
	commands, err := readArchive(path)
	if err != nil {
		return nil, err
	}
	for _, command := range commands {
		comm := strings.Split(command, " ")

		out, err := exec.Command(comm[0], comm[1:]...).CombinedOutput()
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, string(out))
	}
	return outputs, nil
}
