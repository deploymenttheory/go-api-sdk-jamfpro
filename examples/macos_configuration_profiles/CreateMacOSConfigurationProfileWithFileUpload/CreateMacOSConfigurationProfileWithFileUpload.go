package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	maxConcurrentRequestsAllowed = 5 // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

// Helper function to get a pointer from a string
func stringPtr(s string) *string {
	return &s
}

// readPayloadFromFile loads config profile for upload
func readPayloadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:             authConfig.InstanceName,
		DebugMode:                true,
		Logger:                   jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests:    maxConcurrentRequestsAllowed,
		TokenLifespan:            defaultTokenLifespan,
		TokenRefreshBufferPeriod: defaultBufferPeriod,
		ClientID:                 authConfig.ClientID,
		ClientSecret:             authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Set OAuth credentials for the client's HTTP client using the helper function
	creds := map[string]string{
		"clientID":     authConfig.ClientID,
		"clientSecret": authConfig.ClientSecret,
	}
	client.HTTP.SetAuthenticationCredentials(creds)

	// Read the payload from a file
	payloadFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/api-test.mobileconfig"
	payload, err := readPayloadFromFile(payloadFilePath)
	if err != nil {
		log.Fatalf("Error reading payload from file: %v", err)
	}

	// Define a sample macOS Configuration Profile
	profile := &jamfpro.ResponseMacOSConfigurationProfile{
		General: jamfpro.GeneralConfig{
			Name:               "WiFi Test",
			Site:               jamfpro.SiteInfo{Name: "None"},
			Category:           jamfpro.CategoryInfo{Name: "No category assigned"},
			DistributionMethod: "Install Automatically",
			Level:              "computer",
			RedeployOnUpdate:   "Newly Assigned",
			Payloads:           payload,
		},
		Scope: jamfpro.ScopeConfig{},
		SelfService: jamfpro.SelfServiceConfig{
			InstallButtonText: "Install",
		},
	}

	// Call CreateMacOSConfigurationProfile function
	createdProfile, err := client.CreateMacOSConfigurationProfile(profile)
	if err != nil {
		log.Fatalf("Error creating macOS Configuration Profile: %v", err)
	}

	// Print the created profile details
	fmt.Printf("Created macOS Configuration Profile with ID %d and Name %s\n", createdProfile.General.ID, createdProfile.General.Name)
}
