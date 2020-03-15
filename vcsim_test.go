package vcsim_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/matthewcosgrove/vcsim-test-helper"
	"github.com/stretchr/testify/assert"
)

func TestStartVcsimInBackground(t *testing.T) {
	got, err := vcsim.StartVcsimInBackground()
	if err != nil {
		t.Errorf("Error should be nil but was %w\n", err)
	}
	assert.Equal(t, -1, got.ProcessState.ExitCode(), "Process should be running")
	callGovcFind()
	got.Process.Kill()

}

func callGovcFind() {

	cmd := exec.Command("govc", "find", "-l")
	cmd.Env = append(cmd.Env, "GOVC_INSECURE=true")
	cmd.Env = append(cmd.Env, "GOVC_URL=https://user:pass@127.0.0.1:8989/sdk")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
