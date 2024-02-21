package applescript

import (
	"bytes"
	"fmt"
	"os/exec"
	"text/template"
)

// InspectPackageWithAppleScript demonstrates the use of AppleScript to interact with a Cocoa application
// by opening a specified .pkg file on the desktop. It showcases how to access installed files and folders within
// the package by name, using the "Suspicious Package" application for macOS. The function illustrates
// the use of the 'reveal' command to display an installed item within the package in a new tab of the
// "Suspicious Package" user interface, providing a practical example of manipulating GUI elements and
// retrieving file properties through AppleScript from Go.
func InspectPackageWithAppleScript(packageName string) error {
	// Template for the AppleScript, with {{.PackageName}} as a placeholder for the package name
	appleScriptTemplate := `
tell application "Finder"
	set thePackage to (path to desktop as string) & "{{.PackageName}}"
end tell
tell application "Suspicious Package"
	-- tell Suspicious Package to open the package
	set theDocument to (open file thePackage)
	-- find any launchd agent plist directory, by partial POSIX path
	set launchAgents to (get installed item "System/Library/LaunchAgents" of theDocument)
	-- does the package install to the launchd agent directory?
	if exists launchAgents then
		-- examine each launch agent plist in the package
		repeat with anAgent in installed items of launchAgents
			-- get some properties of the plist
			display notification "Found " & (name of anAgent) & " with owner " & (owner of anAgent)
			log (get URL of anAgent)
			-- reveal the plist in a new tab in the Suspicious Package UI
			reveal anAgent
		end repeat
	end if
	-- close the package
	close theDocument
end tell
`

	// Create a new template and parse the AppleScript template text
	tmpl, err := template.New("applescript").Parse(appleScriptTemplate)
	if err != nil {
		return fmt.Errorf("error creating template: %w", err)
	}

	// Execute the template with the packageName parameter to fill in the placeholder
	var appleScript bytes.Buffer
	err = tmpl.Execute(&appleScript, map[string]string{"PackageName": packageName})
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// Execute the filled-in AppleScript using os/exec
	cmd := exec.Command("osascript", "-e", appleScript.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing AppleScript: %w", err)
	}

	// Output from the AppleScript execution, if needed
	fmt.Println("AppleScript Output:", out.String())
	return nil
}
