package pledge

import (
	"os/exec"
	"syscall"
	"testing"
)

func TestMain(t *testing.T) {
	// We need to do the equivalent of "ulimit -c 0" , so that we don't litter the directory with core dumps.
	syscall.Setrlimit(syscall.RLIMIT_CORE, &syscall.Rlimit{0, 0})
}

type test struct {
	cmd    string
	output string
}

// If we call pledge(2) in the test binary, the process will be restricted for all following tests.
// Instead test programs live in testdata/ and we exec separate processes for each test.

var aborts []test = []test{
	{"stdio", ""},
}

var successes []test = []test{
	{"undefined", "Pledged!\n"},
}

func TestAborts(t *testing.T) {
	for _, v := range aborts {
		want := v.output + "signal: abort trap\n"
		out, err := exec.Command("go", "run", "testdata/"+v.cmd+".go").CombinedOutput()
		if string(out) != want || err == nil {
			t.Fatalf("Output is %s with error %v, want %s\n", out, err, want)
		}
	}
}

func TestSuccesses(t *testing.T) {
	for _, v := range successes {
		out, err := exec.Command("go", "run", "testdata/"+v.cmd+".go").CombinedOutput()
		if string(out) != v.output || err != nil {
			t.Fatalf("Output is %s with error %v, want %s\n", out, err, v.output)
		}
	}
}
