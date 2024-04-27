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
	config := CreatePasscodeConfig()

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
		filePath = filepath.Join(filePath, "PasscodePolicy.mobileconfig")
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

// CreatePasscodeConfig initializes and returns a ResourcePasscodeConfigurationProfile with sample data.
func CreatePasscodeConfig() models.ResourcePasscodeConfigurationProfile {
	trueVal := true
	falseVal := false
	descriptions := []models.Description{
		{
			Locale:      "en-US",
			Description: "Must include a symbol and an uppercase letter.",
		},
		// Add more descriptions for other locales as needed
	}

	return models.ResourcePasscodeConfigurationProfile{
		Version: "1.0",
		PayloadContent: []models.PasscodeConfigurationProfileSubsetPayload{
			{
				AllowSimple:         &trueVal,
				ForcePIN:            &trueVal,
				MaxFailedAttempts:   5,
				MaxGracePeriod:      1,
				MaxInactivity:       2,
				MaxPINAgeInDays:     30,
				MinLength:           8,
				PinHistory:          2,
				RequireAlphanumeric: &falseVal,
				PayloadIdentifier:   "com.example.mypasscodepayload",
				PayloadType:         "com.apple.mobiledevice.passwordpolicy",
				PayloadUUID:         "2a8a75e5-d17d-44d5-b062-3cb92161af9f",
				PayloadVersion:      1,
				CustomRegex: &models.CustomRegex{
					PasswordContentDescriptions: descriptions,
					PasswordContentRegex:        `^(?=.*[A-Z])(?=.*[!@#$&*]).*$`,
				},
			},
		},
		PayloadDisplayName: "Passcode",
		PayloadIdentifier:  "com.example.myprofile",
		PayloadType:        "Configuration",
		PayloadUUID:        "e044f50d-ff67-4bcd-9f3f-d7b678091061",
		PayloadVersion:     1,
	}
}
