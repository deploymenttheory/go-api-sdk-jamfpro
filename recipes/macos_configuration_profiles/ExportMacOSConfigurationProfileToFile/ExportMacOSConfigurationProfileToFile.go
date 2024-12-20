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
	exportDir      = "/Users/dafyddwatkins/localtesting/terraform/support_files/macosconfigurationprofiles/imazing/post-jamfpro-upload"
	profileID      = "5498" // Set your desired profile ID here
)

func main() {
	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch single profile by ID
	respProfile, err := client.GetMacOSConfigurationProfileByID(profileID)
	if err != nil {
		log.Fatalf("Failed to fetch profile with ID %s: %v", profileID, err)
	}

	fmt.Printf("Exporting profile: %s (ID: %s)\n", respProfile.General.Name, profileID)

	if err := os.MkdirAll(exportDir, 0750); err != nil {
		log.Fatalf("Failed to create export directory: %v", err)
	}

	xmlData, err := xml.MarshalIndent(respProfile, "", "  ")
	if err != nil {
		log.Fatalf("Failed to convert profile to XML: %v", err)
	}

	payloadsContent := extractPayloads(string(xmlData))
	if payloadsContent == "" {
		log.Fatal("No <payloads> content found in the profile")
	}

	// Remove escaped characters
	reformattedPayloads, err := removeEscapedCharacters(payloadsContent)
	if err != nil {
		log.Fatalf("Failed to reformat payloads: %v", err)
	}

	filename := filepath.Join(exportDir, respProfile.General.Name+".mobileconfig")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(reformattedPayloads); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	fmt.Printf("Successfully exported profile to %s\n", filename)
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
