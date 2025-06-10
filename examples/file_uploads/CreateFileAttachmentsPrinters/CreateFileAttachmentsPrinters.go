package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "./clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example: Upload a file to a computer with ID 2
	err = client.CreateFileAttachment(
		"printers", // resource type / computers / mobiledevices / enrollment profiles / printers / peripherals / policies /  ebooks / mobiledeviceapplicationsicon / mobiledeviceapplicationsipa / diskencryptionconfigurations
		"id",       // idType - can be either 'id' or 'name'.
		"2419",     // identifier
		"/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/icon/UploadIcon/cat.png", // filePath
		false, // forceIpaUpload
	)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Println("File uploaded successfully")
}
