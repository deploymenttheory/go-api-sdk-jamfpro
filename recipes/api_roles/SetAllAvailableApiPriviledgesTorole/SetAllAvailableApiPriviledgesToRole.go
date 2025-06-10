package main

import (
	"fmt"
	"log"
	"strings"

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

	// Get Jamf Pro Version
	versionInfo, err := client.GetJamfProVersion()
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro version: %v", err)
	}

	if versionInfo.Version == nil {
		log.Fatal("Received empty version information")
	}

	// Clean version string by removing any spaces or special characters
	version := strings.TrimSpace(*versionInfo.Version)
	fmt.Printf("Found Jamf Pro version: %s\n", version)

	// Get all available privileges
	privileges, err := client.GetJamfAPIPrivileges()
	if err != nil {
		log.Fatalf("Failed to get API privileges: %v", err)
	}

	if len(privileges.Privileges) == 0 {
		log.Fatal("No privileges found")
	}

	fmt.Printf("Found %d privileges\n", len(privileges.Privileges))

	// Create the role name using the version
	roleName := fmt.Sprintf("all-jamfpro-privileges-%s", version)

	// Create new role with all privileges
	newRole := &jamfpro.ResourceAPIRole{
		DisplayName: roleName,
		Privileges:  privileges.Privileges,
	}

	createdRole, err := client.CreateJamfApiRole(newRole)
	if err != nil {
		log.Fatalf("Failed to create API role: %v", err)
	}

	fmt.Printf("Successfully created role '%s' with ID: %s\n", createdRole.DisplayName, createdRole.ID)
	fmt.Printf("Total privileges assigned: %d\n", len(createdRole.Privileges))
}
