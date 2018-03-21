package build

import (
	"os"
	"os/exec"
	"fmt"
)

func BuildNgrok(targetFile string, goOS OS, goArch Arch) error{
	args := []string{
		"build",
		"-o", 
		targetFile,
	}
	args = append(args, "-tags", "release")
	args = append(args, "v2ray.com/core/_ngrok/main/ngrok")

	for index, value := range args {
		fmt.Printf("args[%d]=%d \n", index, value)
	}

	cmd := exec.Command("go", args...)
	cmd.Env = append(cmd.Env, "GOOS="+string(goOS), "GOARCH="+string(goArch), "CGO_ENABLED=0")
	cmd.Env = append(cmd.Env, os.Environ()...)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		os.Stdout.Write(output)
	}

	return err
}