package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the path to the icon file
	iconPath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/icon/UploadIcon/cat.png"

	response, err := client.UploadIcon(iconPath)
	if err != nil {
		fmt.Println("Error uploading icon:", err)
		return
	}

	// Pretty print the icon details
	iconJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling accounts data: %v", err)
	}
	fmt.Println("Fetched Icon:", string(iconJSON))
}
