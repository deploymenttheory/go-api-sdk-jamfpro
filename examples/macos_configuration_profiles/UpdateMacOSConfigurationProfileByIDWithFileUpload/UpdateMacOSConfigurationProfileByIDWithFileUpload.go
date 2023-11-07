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
	maxConcurrentRequestsAllowed = 5
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func readPayloadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Load payload from file
	payloads, err := readPayloadFromFile("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/api-test.mobileconfig")
	if err != nil {
		log.Fatalf("Failed to read payload: %v", err)
	}

	generalConfig := jamfpro.GeneralConfig{
		Name:               "WiFi Test Updated with sdk",
		Description:        "",
		Site:               jamfpro.SiteInfo{Name: "None"},
		Category:           jamfpro.CategoryInfo{Name: "No category assigned"},
		DistributionMethod: "Install Automatically",
		UserRemovable:      false,
		Level:              "computer",
		RedeployOnUpdate:   "Newly Assigned",
		Payloads:           payloads,
	}

	scopeConfig := jamfpro.ScopeConfig{
		AllComputers: false,
		AllJSSUsers:  false,
	}

	selfServiceConfig := jamfpro.SelfServiceConfig{
		InstallButtonText:           "Install",
		SelfServiceDescription:      "null",
		ForceUsersToViewDescription: false,
		FeatureOnMainPage:           false,
	}

	profile := &jamfpro.ResponseMacOSConfigurationProfile{
		General:     generalConfig,
		Scope:       scopeConfig,
		SelfService: selfServiceConfig,
	}

	// Assuming the ID of the macOS Configuration Profile you want to update is 123
	id := 78

	// Call the UpdateMacOSConfigurationProfileByID function
	response, err := client.UpdateMacOSConfigurationProfileByID(id, profile)
	if err != nil {
		log.Fatalf("Failed to update macOS Configuration Profile: %v", err)
	}

	fmt.Printf("Profile updated: %+v\n", response)
}
