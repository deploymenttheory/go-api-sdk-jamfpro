package main

import (
	"fmt"
	"log"

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

	// Define the database password
	databasePassword := "your-secure-database-password"

	// Call InitializeDatabaseConnection function
	err = client.InitializeDatabaseConnection(databasePassword)
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}

	fmt.Println("Successfully initialized database connection.")
}
