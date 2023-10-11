package main

import (
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:          authConfig.InstanceName,
		DebugMode:             true,
		Logger:                jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests: maxConcurrentRequestsAllowed,
		TokenLifespan:         defaultTokenLifespan,
		BufferPeriod:          defaultBufferPeriod,
		ClientID:              authConfig.ClientID,
		ClientSecret:          authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client := jamfpro.NewClient(config)

	// 2. Construct the macOS Configuration Profile data
	payloads := `&lt;?xml version="1.0" encoding="UTF-8"?&gt;&lt;!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"&gt;
	&lt;plist version="1"&gt;&lt;dict&gt;&lt;key&gt;PayloadUUID&lt;/key&gt;&lt;string&gt;3665BB04-2B24-4CF6-A3FB-BF1B2221CBA5&lt;/string&gt;&lt;key&gt;PayloadType&lt;/key&gt;&lt;string&gt;Configuration&lt;/string&gt;&lt;key&gt;PayloadOrganization&lt;/key&gt;&lt;string&gt;Jamf&lt;/string&gt;&lt;key&gt;PayloadIdentifier&lt;/key&gt;&lt;string&gt;3665BB04-2B24-4CF6-A3FB-BF1B2221CBA5&lt;/string&gt;&lt;key&gt;PayloadDisplayName&lt;/key&gt;&lt;string&gt;WiFi Test&lt;/string&gt;&lt;key&gt;PayloadDescription&lt;/key&gt;&lt;string/&gt;&lt;key&gt;PayloadVersion&lt;/key&gt;&lt;integer&gt;1&lt;/integer&gt;&lt;key&gt;PayloadEnabled&lt;/key&gt;&lt;true/&gt;&lt;key&gt;PayloadRemovalDisallowed&lt;/key&gt;&lt;true/&gt;&lt;key&gt;PayloadScope&lt;/key&gt;&lt;string&gt;System&lt;/string&gt;&lt;key&gt;PayloadContent&lt;/key&gt;&lt;array&gt;&lt;dict&gt;&lt;key&gt;PayloadUUID&lt;/key&gt;&lt;string&gt;646F7DC3-54EF-42AF-92BA-6626DA39E28F&lt;/string&gt;&lt;key&gt;PayloadType&lt;/key&gt;&lt;string&gt;com.apple.wifi.managed&lt;/string&gt;&lt;key&gt;PayloadOrganization&lt;/key&gt;&lt;string&gt;Jamf&lt;/string&gt;&lt;key&gt;PayloadIdentifier&lt;/key&gt;&lt;string&gt;646F7DC3-54EF-42AF-92BA-6626DA39E28F&lt;/string&gt;&lt;key&gt;PayloadDisplayName&lt;/key&gt;&lt;string&gt;WiFi&lt;/string&gt;&lt;key&gt;PayloadDescription&lt;/key&gt;&lt;string/&gt;&lt;key&gt;PayloadVersion&lt;/key&gt;&lt;integer&gt;1&lt;/integer&gt;&lt;key&gt;PayloadEnabled&lt;/key&gt;&lt;true/&gt;&lt;key&gt;HIDDEN_NETWORK&lt;/key&gt;&lt;false/&gt;&lt;key&gt;Password&lt;/key&gt;&lt;string&gt;jamf&lt;/string&gt;&lt;key&gt;EncryptionType&lt;/key&gt;&lt;string&gt;WPA&lt;/string&gt;&lt;key&gt;AutoJoin&lt;/key&gt;&lt;true/&gt;&lt;key&gt;CaptiveBypass&lt;/key&gt;&lt;false/&gt;&lt;key&gt;ProxyType&lt;/key&gt;&lt;string&gt;None&lt;/string&gt;&lt;key&gt;SetupModes&lt;/key&gt;&lt;array/&gt;&lt;key&gt;SSID_STR&lt;/key&gt;&lt;string&gt;jamf&lt;/string&gt;&lt;key&gt;Interface&lt;/key&gt;&lt;string&gt;BuiltInWireless&lt;/string&gt;&lt;/dict&gt;&lt;/array&gt;&lt;/dict&gt;&lt;/plist&gt;` // Your XML payload here

	// General profile data
	generalConfig := jamfpro.GeneralConfig{
		Name:               "WiFi Test",
		Description:        "",
		Site:               jamfpro.SiteInfo{Name: "None"},
		Category:           jamfpro.CategoryInfo{Name: "No category assigned"},
		DistributionMethod: "Install Automatically",
		UserRemovable:      false,
		Level:              "computer",
		RedeployOnUpdate:   "Newly Assigned",
		Payloads:           payloads,
	}

	// Scope data
	scopeConfig := jamfpro.ScopeConfig{
		AllComputers: false,
		AllJSSUsers:  false,
	}

	// SelfService data
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

	// Call the CreateMacOSConfigurationProfile function
	response, err := client.CreateMacOSConfigurationProfile(profile)
	if err != nil {
		log.Fatalf("Failed to create macOS Configuration Profile: %v", err)
	}

	// 4. Print the response or handle it as needed
	fmt.Printf("Profile created: %+v\n", response)
}
