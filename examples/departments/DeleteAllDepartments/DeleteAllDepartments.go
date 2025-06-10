package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch all departments
	// For more information on how to add parameters to this request, see docs/url_queries.md
	departments, err := client.GetDepartments(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching departments: %v", err)
	}

	fmt.Println("Departments fetched. Starting deletion process:")

	// Iterate over each department and delete
	for _, department := range departments.Results {
		fmt.Printf("Deleting department ID: %s, Name: %s\n", department.ID, department.Name)

		err = client.DeleteDepartmentByID(department.ID)
		if err != nil {
			log.Printf("Error deleting department ID %s: %v\n", department.ID, err)
			continue // Move to the next department if there's an error
		}

		fmt.Printf("Department ID %s deleted successfully.\n", department.ID)
	}

	fmt.Println("Department deletion process completed.")
}
