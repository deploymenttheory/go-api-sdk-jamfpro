package modules

import (
	"fmt"
	"os/exec"
	"strings"
)

// Function to get the hardware UUID of the Mac
func GetHardwareUUIDFromSystemProfiler() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute system_profiler: %v", err)
	}

	// Parse output to find Hardware UUID
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Hardware UUID") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("hardware UUID not found in system_profiler output")
}

// Function to get the Mac's serial number
func GetSerialNumberFromSystemProfiler() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute system_profiler: %v", err)
	}

	// Parse output to find Serial Number
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("serial number not found in system_profiler output")
}
