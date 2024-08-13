package main

import (
	"encoding/json"
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

	// Define a sort filter (you can modify this as needed)
	sortFilter := "sort=name:asc"

	// Call GetJamfProtectPlans function
	plans, err := client.GetJamfProtectPlans(sortFilter)
	if err != nil {
		log.Fatalf("Error fetching Jamf Protect plans: %v", err)
	}

	// Print the total count of plans
	fmt.Printf("Total Jamf Protect plans: %d\n\n", plans.TotalCount)

	// Print details of each plan
	for _, plan := range plans.Results {
		fmt.Printf("Plan Name: %s\n", plan.Name)
		fmt.Printf("UUID: %s\n", plan.UUID)
		fmt.Printf("ID: %s\n", plan.ID)
		fmt.Printf("Description: %s\n", plan.Description)
		fmt.Printf("Profile ID: %d\n", plan.ProfileID)
		fmt.Printf("Profile Name: %s\n", plan.ProfileName)
		fmt.Printf("Scope Description: %s\n", plan.ScopeDescription)
		fmt.Println("--------------------")
	}

	// Optionally, you can also print the entire response as JSON
	plansJSON, err := json.MarshalIndent(plans, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect plans to JSON: %v", err)
	}
	fmt.Println("Jamf Protect Plans (JSON):")
	fmt.Println(string(plansJSON))
}
