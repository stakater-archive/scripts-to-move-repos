package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunCommandVerbose runs the command displaying its output
func RunCommandVerbose(name string, args ...string) error {
	e := exec.Command(name, args...)
	err := e.Run()
	if err != nil {
		fmt.Printf("Error: Command failed  %s %s\n", name, strings.Join(args, " "))
	}
	return err
}
