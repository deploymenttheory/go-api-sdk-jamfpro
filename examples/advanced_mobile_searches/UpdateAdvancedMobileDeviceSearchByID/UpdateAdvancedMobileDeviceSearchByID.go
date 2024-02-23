package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the advanced mobile device search details
	updatedSearch := &jamfpro.ResourceAdvancedMobileDeviceSearch{
		Name:   "jamf api sdk advanced search",
		ViewAs: "Standard Web Page",
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 1,
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "more than x days ago",
					Value:        "7",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.SharedAdvancedSearchContainerDisplayField{
			{
				DisplayField: []jamfpro.SharedAdvancedSearchSubsetDisplayField{
					{
						Name: "AirPlay Password",
					},
					{
						Name: "App Analytics Enabled",
					},
					// Add more display fields as needed
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	searchID := 14

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(updatedSearch, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Create the advanced mobile device search
	updatedSearchResp, err := client.UpdateAdvancedMobileDeviceSearchByID(searchID, updatedSearch)
	if err != nil {
		fmt.Println("Error updating advanced mobile device search:", err)
		return
	}

	// Print the created advanced mobile device search details
	createdSearchXML, err := xml.MarshalIndent(updatedSearchResp, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling created search to XML:", err)
		return
	}
	fmt.Printf("updated Advanced mobile device Search:\n%s\n", string(createdSearchXML))
}
