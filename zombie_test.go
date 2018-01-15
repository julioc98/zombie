package zombie

import (
	"testing"
)

func TestReadArchive(t *testing.T) {
	var in, out = "config", []string{"echo 'oi eu sou' oi 'oi sou goku'", "echo 2"}
	arc, _ := readArchive(in)

	for i, a := range arc {
		if a != out[i] {
			t.Errorf("readArchive(%v) = %v, want %v", in, arc, out)
		}
	}

	// Error case
	var inErr = "configx"
	if _, err := readArchive(inErr); err == nil {
		t.Errorf("readArchive(%v) = %v", inErr, err)
	}
}

func TestCleanCommand(t *testing.T) {
	var (
		in  = "commit -m \"message with space\" hi -m 'another message'"
		out = []string{"commit", "-m", "\"message with space\"", "hi", "-m", "'another message'"}
	)
	comm, _ := cleanCommand(in)
	for i, c := range comm {
		if c != out[i] {
			t.Errorf("readArchive(%v) = %v, want %v", in, comm, out)
		}
	}

}

func TestExecCommandPath(t *testing.T) {
	var in, out = "config", []string{"'oi eu sou' oi 'oi sou goku'\n", "2\n"}
	exec, _ := ExecCommandPath(in)

	for i, e := range exec {
		if e != out[i] {
			t.Errorf("readArchive(%v) = %v, want %v", in, exec, out)
		}
	}

	// Error case one
	var inErr = "configx"
	if _, err := ExecCommandPath(inErr); err == nil {
		t.Errorf("ExecCommandPath(%v) = %v", inErr, err)

	}
	// Error case two
	var inArch = "config_empy"
	if _, err := ExecCommandPath(inArch); err == nil {
		t.Errorf("ExecCommandPath(%v) = %v", inArch, err)
	}
}
