package abigen

import (
	"bytes"
	"os/exec"

	"github.com/pkg/errors"
)

func Run(
	abi string,
	typ string,
	pkg string,
	output string,
	bin string,
) error {
	args := []string{
		"-abi", abi,
		"-out", output,
		"-type", typ,
		"-pkg", pkg,
		"-bin", bin,
	}

	cmd := exec.Command("abigen", args...)

	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return errors.Errorf("Failed to generate bindings: %s, stderr: %s", err, stdErr.String())
	}

	return nil
}
