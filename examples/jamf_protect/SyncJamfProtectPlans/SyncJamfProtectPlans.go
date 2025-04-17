package main

import (
	"fmt"
	"log"
	"net/url"

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

	// Call SyncJamfProtectPlans function
	err = client.SyncJamfProtectPlans()
	if err != nil {
		log.Fatalf("Error syncing Jamf Protect plans: %v", err)
	}

	// If we reach this point, the sync was successful
	fmt.Println("Jamf Protect plans synced successfully")

	// Optional: You could add additional operations here, such as fetching and displaying the synced plans
	// For more information on how to add parameters to this request, see docs/url_queries.md
	// For example:
	plans, err := client.GetJamfProtectPlans(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching Jamf Protect plans: %v", err)
	}

	fmt.Printf("Synced %d Jamf Protect plans\n", plans.TotalCount)
	for _, plan := range plans.Results {
		fmt.Printf("Plan: %s (ID: %s)\n", plan.Name, plan.ID)
	}
}
