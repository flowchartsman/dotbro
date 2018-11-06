// +build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var runCmd string

func prepInit(outputer Outputer) error {
	var err error
	runCmd, err = exec.LookPath("bash")
	if err != nil {
		runCmd, err = exec.LookPath("sh")
		if err != nil {
			return fmt.Errorf("unable to find 'bash' or 'sh', cannot run init commands: %s", err)
		}
		outputer.OutWarn("'bash' missing, using 'sh'")
	}
	return nil
}

func runCommand(command string, outputer Outputer) error {
	cmd := exec.Command(runCmd, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
