package zombie

import (
	"reflect"
	"testing"
)

func TestReadArchive(t *testing.T) {
	var in, out = "config", []string{"echo oi\n", "2\n"}
	if x, _ := readArchive(in); reflect.DeepEqual(x, out) {
		t.Errorf("readArchive(%v) = %v, want %v", in, x, out)
	}
}
func TestExecCommandPath(t *testing.T) {
	var in, out = "config", []string{"oi", "2"}
	if x, _ := ExecCommandPath(in); reflect.DeepEqual(x, out) {
		t.Errorf("ExecCommandPath(%v) = %v, want %v", in, x, out)
	}
}
