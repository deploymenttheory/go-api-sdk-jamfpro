package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the path to the image file
	imagePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/examples/self_service_branding_images/UploadSelfServiceBrandingImage/cat.png"

	response, err := client.UploadSelfServiceBrandingImage(imagePath)
	if err != nil {
		fmt.Println("Error uploading image:", err)
		return
	}

	// Pretty print the image details
	imageJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling image data: %v", err)
	}
	fmt.Println("Uploaded Image:", string(imageJSON))
}
