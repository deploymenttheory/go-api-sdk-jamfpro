package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "./clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	packageName := "Firefox 133.0.3.pkg"

	response, err := client.GetPackageByName(packageName)
	if err != nil {
		fmt.Println("Error fetching package by name:", err)
		return
	}

	packageJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling package data: %v", err)
	}
	fmt.Println("Obtained package Details:\n", string(packageJSON))
}
