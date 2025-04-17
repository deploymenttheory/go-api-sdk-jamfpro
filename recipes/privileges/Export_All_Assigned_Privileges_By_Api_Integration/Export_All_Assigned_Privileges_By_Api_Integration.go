package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch all API integrations
	// For more information on how to add parameters to this request, see docs/url_queries.md
	integrations, err := client.GetApiIntegrations(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching API integrations: %v", err)
	}

	// Create a folder for the CSV file
	folderPath := "jamf_pro_exports"
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatalf("Error creating folder: %v", err)
	}

	// Create a CSV file in the new folder
	filePath := filepath.Join(folderPath, "api_integrations_export.csv")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Integration ID", "Display Name", "Enabled", "Access Token Lifetime", "App Type", "Client ID", "Authorization Scopes", "API Privileges"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Error writing CSV header: %v", err)
	}

	// Process each integration
	for _, integration := range integrations.Results {
		// Fetch API roles for this integration
		// For more information on how to add parameters to this request, see docs/url_queries.md
		roles, err := client.GetJamfAPIRoles(url.Values{})
		if err != nil {
			log.Printf("Error fetching API roles for integration %d: %v", integration.ID, err)
			continue
		}

		// Collect all privileges for this integration
		var privileges []string
		for _, role := range roles.Results {
			privileges = append(privileges, role.Privileges...)
		}

		// Remove duplicates from privileges
		uniquePrivileges := removeDuplicates(privileges)

		// Create CSV record
		record := []string{
			fmt.Sprintf("%d", integration.ID),
			integration.DisplayName,
			fmt.Sprintf("%t", integration.Enabled),
			fmt.Sprintf("%d", integration.AccessTokenLifetimeSeconds),
			integration.AppType,
			integration.ClientID,
			strings.Join(integration.AuthorizationScopes, ", "),
			strings.Join(uniquePrivileges, ", "),
		}

		if err := writer.Write(record); err != nil {
			log.Printf("Error writing record for integration %d: %v", integration.ID, err)
		}
	}

	fmt.Printf("Export completed successfully. File saved as %s\n", filePath)
}

func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
