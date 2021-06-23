package devicebattery

import (
	"os/exec"
	"strings"
)

// GetDevices gets the devices of a particular type
func GetDevices(deviceType []string) (string, error) {
	defaultArgs := []string{"-json"}
	sysProfilerArgs := []string{}
	sysProfilerArgs = append(defaultArgs, deviceType...)
	cmd := exec.Command("system_profiler", sysProfilerArgs...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
