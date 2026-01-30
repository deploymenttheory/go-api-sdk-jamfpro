package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetLdapServersV1 to retrieve LDAP and Cloud Identity Provider servers
	servers, err := client.GetLdapServersV1()
	if err != nil {
		log.Fatalf("Error fetching LDAP servers: %v", err)
	}

	// Pretty print the server list
	serversJSON, err := json.MarshalIndent(servers, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling LDAP servers data: %v", err)
	}

	fmt.Println("Fetched LDAP Servers:\n", string(serversJSON))
}
