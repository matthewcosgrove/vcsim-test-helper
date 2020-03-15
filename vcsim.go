package vcsim

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func StartVcsimInBackground() (*exec.Cmd, error) {
	command := "vcsim"

	path, err := exec.LookPath(command)
	if err != nil {
		return nil, fmt.Errorf("'%s' executable not found!\n", command)
	}
	fmt.Printf("Using %s in %s\n", command, path)

	cmd := exec.Command(command)
	return startInBackground(cmd)

}

func startInBackground(cmd *exec.Cmd) (*exec.Cmd, error) {
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	command := cmd.Path
	fmt.Printf("Ensuring %q running with PID: %d\n", command, cmd.Process.Pid)
	if err != nil {
		return nil, fmt.Errorf("Have you installed %q?: %w", command, err)
	}

	go func() {
		err := cmd.Wait()
		if err.Error() == "exit status 2" {
			fmt.Printf("PID: %d\n", cmd.Process.Pid)
			panic(fmt.Errorf("%q cannot start, is it already running? Check pids in ps aux output above\n%s\n", command, err))
		}
		fmt.Printf("Expected error occurred: %s\n", err)
	}()
	time.Sleep(2 * time.Second)

	return cmd, nil
}
