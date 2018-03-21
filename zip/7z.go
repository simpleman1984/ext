package zip

import (
	"os"
	"os/exec"
)

func sevenZipFolder(folder string, targetFile string) error {
	args := []string{
		"a", "-tzip", "-mx=9", targetFile, folder,
	}

	cmd := exec.Command("7z", args...)
	cmd.Env = append(cmd.Env, os.Environ()...)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		os.Stdout.Write(output)
	}

	return err
}
