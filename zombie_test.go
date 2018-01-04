package zombie

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadArchive(t *testing.T) {
	var in, out = "config", []string{"echo 'oi eu sou' oi 'oi sou goku'", "echo 2"}
	if x, _ := readArchive(in); !reflect.DeepEqual(x, out) {
		t.Errorf("readArchive(%v) = %v, want %v", in, x, out)
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
	if x, _ := cleanCommand(in); !reflect.DeepEqual(x, out) {
		t.Errorf("cleanCommand(%v) = %v, want %v || %v --- %s", in, x, out, strings.Join(out, "_"), strings.Join(x, "_"))
	}
}

func TestExecCommandPath(t *testing.T) {
	var in, out = "config", []string{"'oi eu sou' oi 'oi sou goku'\n", "2\n"}
	if x, _ := ExecCommandPath(in); !reflect.DeepEqual(x, out) {
		t.Errorf("ExecCommandPath(%v) = %v, want %v", in, x, out)
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
