package modules

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	// Path to Jamf preferences file
	jamfBinaryPlist = "/Library/Preferences/com.jamfsoftware.jamf.plist"

	// Path to Jamf binary
	jamfBinaryFilePath = "/usr/local/bin/jamf"
)

// Function to retrieve the Jamf Pro (JSS) URL from preferences
func GetJamfProURL() (string, error) {
	// Check if file exists
	if _, err := os.Stat(jamfBinaryPlist); os.IsNotExist(err) {
		return "", fmt.Errorf("%s not found", jamfBinaryPlist)
	}

	// Use defaults command to read plist
	cmd := exec.Command("defaults", "read", jamfBinaryPlist, "jss_url")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to read JSS URL from preferences: %v", err)
	}

	// Trim whitespace and trailing slashes
	jssURL := strings.TrimSpace(string(output))
	jssURL = strings.TrimRight(jssURL, "/")

	return jssURL, nil
}

// Function to check connectivity to the Jamf Pro server
func CheckJamfProConnection() error {
	cmd := exec.Command(jamfBinaryFilePath, "-checkjssconnection", "-retry", "10")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("JSS connection not active: %v", err)
	}

	fmt.Println("JSS connection active!")
	return nil
}
