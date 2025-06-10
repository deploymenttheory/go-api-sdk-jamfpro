package main

import (
	"fmt"
	"log"

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

	// 2. Construct the macOS Configuration Profile data
	payloads := `&lt;?xml version="1.0" encoding="UTF-8"?&gt;&lt;!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"&gt;
	&lt;plist version="1"&gt;&lt;dict&gt;&lt;key&gt;PayloadUUID&lt;/key&gt;&lt;string&gt;3665BB04-2B24-4CF6-A3FB-BF1B2221CBA5&lt;/string&gt;&lt;key&gt;PayloadType&lt;/key&gt;&lt;string&gt;Configuration&lt;/string&gt;&lt;key&gt;PayloadOrganization&lt;/key&gt;&lt;string&gt;Jamf&lt;/string&gt;&lt;key&gt;PayloadIdentifier&lt;/key&gt;&lt;string&gt;3665BB04-2B24-4CF6-A3FB-BF1B2221CBA5&lt;/string&gt;&lt;key&gt;PayloadDisplayName&lt;/key&gt;&lt;string&gt;WiFi Test&lt;/string&gt;&lt;key&gt;PayloadDescription&lt;/key&gt;&lt;string/&gt;&lt;key&gt;PayloadVersion&lt;/key&gt;&lt;integer&gt;1&lt;/integer&gt;&lt;key&gt;PayloadEnabled&lt;/key&gt;&lt;true/&gt;&lt;key&gt;PayloadRemovalDisallowed&lt;/key&gt;&lt;true/&gt;&lt;key&gt;PayloadScope&lt;/key&gt;&lt;string&gt;System&lt;/string&gt;&lt;key&gt;PayloadContent&lt;/key&gt;&lt;array&gt;&lt;dict&gt;&lt;key&gt;PayloadUUID&lt;/key&gt;&lt;string&gt;646F7DC3-54EF-42AF-92BA-6626DA39E28F&lt;/string&gt;&lt;key&gt;PayloadType&lt;/key&gt;&lt;string&gt;com.apple.wifi.managed&lt;/string&gt;&lt;key&gt;PayloadOrganization&lt;/key&gt;&lt;string&gt;Jamf&lt;/string&gt;&lt;key&gt;PayloadIdentifier&lt;/key&gt;&lt;string&gt;646F7DC3-54EF-42AF-92BA-6626DA39E28F&lt;/string&gt;&lt;key&gt;PayloadDisplayName&lt;/key&gt;&lt;string&gt;WiFi&lt;/string&gt;&lt;key&gt;PayloadDescription&lt;/key&gt;&lt;string/&gt;&lt;key&gt;PayloadVersion&lt;/key&gt;&lt;integer&gt;1&lt;/integer&gt;&lt;key&gt;PayloadEnabled&lt;/key&gt;&lt;true/&gt;&lt;key&gt;HIDDEN_NETWORK&lt;/key&gt;&lt;false/&gt;&lt;key&gt;Password&lt;/key&gt;&lt;string&gt;jamf&lt;/string&gt;&lt;key&gt;EncryptionType&lt;/key&gt;&lt;string&gt;WPA&lt;/string&gt;&lt;key&gt;AutoJoin&lt;/key&gt;&lt;true/&gt;&lt;key&gt;CaptiveBypass&lt;/key&gt;&lt;false/&gt;&lt;key&gt;ProxyType&lt;/key&gt;&lt;string&gt;None&lt;/string&gt;&lt;key&gt;SetupModes&lt;/key&gt;&lt;array/&gt;&lt;key&gt;SSID_STR&lt;/key&gt;&lt;string&gt;jamf&lt;/string&gt;&lt;key&gt;Interface&lt;/key&gt;&lt;string&gt;BuiltInWireless&lt;/string&gt;&lt;/dict&gt;&lt;/array&gt;&lt;/dict&gt;&lt;/plist&gt;` // Your XML payload here

	// General profile data
	// Define the macOS Configuration Profile as per the given XML structure
	profile := jamfpro.ResourceMacOSConfigurationProfile{
		General: jamfpro.MacOSConfigurationProfileSubsetGeneral{
			Name:               "WiFi Test",
			Description:        "",
			Site:               &jamfpro.SharedResourceSite{ID: -1, Name: "None"},                     // Optional, the Create fuction will set default values if no site is set
			Category:           &jamfpro.SharedResourceCategory{ID: -1, Name: "No category assigned"}, // Optional, the Create fuction will set default values if no category is set
			DistributionMethod: "Install Automatically",
			UserRemovable:      false,
			Level:              "computer",
			RedeployOnUpdate:   "Newly Assigned",
			Payloads:           payloads,
		},
		Scope: jamfpro.MacOSConfigurationProfileSubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.MacOSConfigurationProfileSubsetSelfService{
			InstallButtonText:           "Install",
			SelfServiceDescription:      "null",
			ForceUsersToViewDescription: false,
			// Add other fields as per the XML example
		},
	}

	// Call the CreateMacOSConfigurationProfile function
	createdProfileID, err := client.CreateMacOSConfigurationProfile(&profile)
	if err != nil {
		fmt.Println("Error creating macOS Configuration Profile:", err)
		return
	}

	// Print the ID of the created profile
	fmt.Printf("Successfully created macOS Configuration Profile with ID: %d\n", createdProfileID)

}
