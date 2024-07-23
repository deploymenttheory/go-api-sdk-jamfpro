package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch all computer extension attributes
	extAtts, err := client.GetComputerExtensionAttributes()
	if err != nil {
		log.Fatalf("Error fetching computer extension attributes: %v", err)
	}

	fmt.Println("computer extension attributes fetched. Starting deletion process:")

	// Iterate over each computer extension attribute and delete
	for _, extAtt := range extAtts.Results {
		fmt.Printf("Deleting computer extension attribute ID: %d, Name: %s\n", extAtt.ID, extAtt.Name)

		err = client.DeleteComputerExtensionAttributeByID(strconv.Itoa(extAtt.ID))
		if err != nil {
			log.Printf("Error deleting computer extension attribute ID %d: %v\n", extAtt.ID, err)
			continue // Move to the next computer extension attribute if there's an error
		}

		fmt.Printf("computer extension attribute ID %d deleted successfully.\n", extAtt.ID)
	}

	fmt.Println("computer extension attribute deletion process completed.")

}
