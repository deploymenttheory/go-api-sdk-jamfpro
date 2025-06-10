package main

import (
	"encoding/xml"
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

	// Call GetScripts function
	// For more information on how to add parameters to this request, see docs/url_queries.md
	scripts, err := client.GetScripts(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching scripts: %v", err)
	}

	// Pretty print the scripts in XML
	scriptsXML, err := xml.MarshalIndent(scripts, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling scripts data: %v", err)
	}
	fmt.Println("Fetched Scripts:\n", string(scriptsXML))
}
