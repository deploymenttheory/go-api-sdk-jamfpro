package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/apple/devicemanagementprofiles/models"
	"howett.net/plist"
)

func main() {
	config := CreateSecurityPreferencesConfig()

	// Serialize the configuration to plist in XML format
	output, err := plist.MarshalIndent(config, plist.XMLFormat, "    ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	// marshalled XML
	fullXML := string(output)

	// Rest of your file handling code...
	fmt.Print("Enter the file path to save the .mobileconfig (e.g., /path/to/file.mobileconfig): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if stat, err := os.Stat(filePath); err == nil && stat.IsDir() {
		filePath = filepath.Join(filePath, "SecurityPreferences.mobileconfig")
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(fullXML))
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}

	fmt.Println("File successfully written to", filePath)
}

// CreateSecurityPreferencesConfig initializes and returns a ResourceSecurityPreferencesConfigurationProfile with predefined values.
func CreateSecurityPreferencesConfig() models.ResourceSecurityPreferencesConfigurationProfile {
	trueVal := true
	falseVal := false
	return models.ResourceSecurityPreferencesConfigurationProfile{
		Version: "1.0",
		PayloadContent: []models.SecurityPreferencesConfigurationProfileSubsetPayload{
			{
				DontAllowFireWallUI:      &trueVal,
				DontAllowLockMessageUI:   &falseVal,
				DontAllowPasswordResetUI: &falseVal,
				PayloadIdentifier:        "com.example.mysecuritypreferencespayload",
				PayloadType:              "com.apple.preference.security",
				PayloadUUID:              "d99bb019-a61d-447f-8fed-8f223cc56be3",
				PayloadVersion:           1,
			},
		},
		PayloadDisplayName: "Security Preferences",
		PayloadIdentifier:  "com.example.myprofile",
		PayloadType:        "Configuration",
		PayloadUUID:        "b44b6a04-6527-4333-87e5-46422e8a5844",
		PayloadVersion:     1,
	}
}
