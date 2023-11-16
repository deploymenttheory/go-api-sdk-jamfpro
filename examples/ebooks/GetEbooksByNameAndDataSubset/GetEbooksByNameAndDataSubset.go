package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ebookName := "iPhone User Guide for iOS 10.3" // Replace with the desired ebook name
	subset := "General"                           // Replace with "General", "Scope", or "SelfService"

	ebook, err := client.GetEbooksByNameAndDataSubset(ebookName, subset)
	if err != nil {
		log.Fatalf("Error fetching ebook by Name and Subset: %v", err)
	}

	ebookXML, err := xml.MarshalIndent(ebook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling ebook data: %v", err)
	}
	fmt.Println("Fetched Ebook Subset:\n", string(ebookXML))
}
