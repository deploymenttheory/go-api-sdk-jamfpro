package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"howett.net/plist"
)

// Global configuration variables
var (
	configFilePath = "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	exportDir      = "/Users/dafyddwatkins/localtesting/jamfpro"
)

func main() {
	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch the list of mobile device configuration profiles
	profilesList, err := client.GetMobileDeviceConfigurationProfiles()
	if err != nil {
		log.Fatalf("Failed to fetch Mobile Device Configuration Profiles: %v", err)
	}

	fmt.Println("Found the following Mobile Device Configuration Profiles:")
	for _, profile := range profilesList.ConfigurationProfiles {
		fmt.Printf("ID: %d, Name: %s\n", profile.ID, profile.Name)
	}
	fmt.Println("These profiles will be exported.")

	if err := os.MkdirAll(exportDir, 0750); err != nil {
		log.Fatalf("Failed to create export directory: %v", err)
	}

	for _, profile := range profilesList.ConfigurationProfiles {
		respProfile, err := client.GetMobileDeviceConfigurationProfileByID(profile.ID)
		if err != nil {
			log.Printf("Failed to fetch profile with ID %d: %v", profile.ID, err)
			continue
		}

		xmlData, err := xml.MarshalIndent(respProfile, "", "  ")
		if err != nil {
			log.Printf("Failed to convert profile with ID %d to XML: %v", profile.ID, err)
			continue
		}

		payloadsContent := extractPayloads(string(xmlData))
		if payloadsContent == "" {
			log.Printf("No <payloads> content found for profile ID %d", profile.ID)
			continue
		}

		// Remove escaped characters
		reformattedPayloads, err := removeEscapedCharacters(payloadsContent)
		if err != nil {
			log.Printf("Failed to reformat payloads for profile with ID %d: %v", profile.ID, err)
			continue
		}

		filename := filepath.Join(exportDir, profile.Name+".mobileconfig")
		file, err := os.Create(filename)
		if err != nil {
			log.Printf("Failed to create file for profile with ID %d: %v", profile.ID, err)
			continue
		}
		defer file.Close()

		if _, err := file.WriteString(reformattedPayloads); err != nil {
			log.Printf("Failed to write to file for profile with ID %d: %v", profile.ID, err)
			continue
		}

		fmt.Printf("Exported profile with ID %d to %s\n", profile.ID, filename)
	}

	fmt.Println("Export completed!")
}

// extractPayloads extracts the payloads from the XML data
func extractPayloads(xmlData string) string {
	startTag := "<payloads>"
	endTag := "</payloads>"
	startIndex := strings.Index(xmlData, startTag)
	endIndex := strings.Index(xmlData, endTag)

	if startIndex == -1 || endIndex == -1 {
		return ""
	}

	return xmlData[startIndex+len(startTag) : endIndex]
}

// removeEscapedCharacters removes escaped characters from a plist / .mobileconfig file
func removeEscapedCharacters(plistContent string) (string, error) {
	data, err := decodePlist([]byte(plistContent))
	if err != nil {
		return "", err
	}

	newPlist, err := plist.MarshalIndent(data, plist.XMLFormat, "\t")
	if err != nil {
		return "", fmt.Errorf("error marshaling plist: %v", err)
	}

	return string(newPlist), nil
}

// decodePlist decodes the plist content and returns the data
func decodePlist(content []byte) (interface{}, error) {
	// Check if content needs unescaping
	if bytes.Contains(content, []byte("&lt;")) {
		content = []byte(html.UnescapeString(string(content)))
	}

	decoder := plist.NewDecoder(bytes.NewReader(content))
	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding plist: %v", err)
	}
	return data, nil
}
