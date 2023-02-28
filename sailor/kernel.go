package sailor

import (
	"fmt"
	"os"
	"os/exec"
)

const sysctlConfigPath = "/etc/sysctl.conf"

func Set(key, value string) error {
	param := fmt.Sprintf("%s=%s", key, value)

	outfile, err := os.Open(sysctlConfigPath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	cmd := exec.Command("sysctl", "-w", param)
	cmd.Stdout = outfile
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
