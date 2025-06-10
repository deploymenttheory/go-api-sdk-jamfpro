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

	// Fetch all printers
	printers, err := client.GetPrinters()
	if err != nil {
		log.Fatalf("Error fetching printers: %v", err)
	}

	fmt.Println("Printers fetched. Starting deletion process:")

	// Iterate over each printer and delete
	for _, printer := range printers.Printer {
		fmt.Printf("deleting printer ID: %d, Name: %s\n", printer.ID, printer.Name)

		err = client.DeletePrinterByID(strconv.Itoa(printer.ID))
		if err != nil {
			log.Printf("error deleting printer ID %d: %v\n", printer.ID, err)
			continue // Move to the next printer if there's an error
		}

		fmt.Printf("printer ID %d deleted successfully.\n", printer.ID)
	}

	fmt.Println("Printer deletion process completed.")
}
