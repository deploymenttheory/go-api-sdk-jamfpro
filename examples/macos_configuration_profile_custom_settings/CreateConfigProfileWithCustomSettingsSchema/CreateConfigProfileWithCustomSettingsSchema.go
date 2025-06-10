package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create the Plist content as a string
	plistContent := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>network_aware</key>
    <true/>
    <key>Install Messages</key>
    <dict>
      <key>notify_installed</key>
      <true/>
    </dict>
    <key>Pending Install Messages</key>
    <dict>
      <key>notify_pending_install</key>
      <true/>
    </dict>
    <key>Pending Uninstall Messages</key>
    <dict>
      <key>notify_pending_uninstall</key>
      <false/>
    </dict>
    <key>Uninstall Messages</key>
    <dict>
      <key>notify_uninstalled</key>
      <true/>
    </dict>
    <key>Up-to-date Messages</key>
    <dict>
      <key>notify_up_to_date</key>
      <true/>
    </dict>
  </dict>
</plist>`

	// Create the JSON Schema content as a string (truncated for brevity)
	jsonSchema := `{"title":"Jamf Auto Update (com.jamf.autoupdate)","description":"Settings for Jamf Auto Update 1.3+.","__version":"1.3","type":"object","links":[{"rel":"More information","href":"https://support.datajar.co.uk/hc/en-us/articles/18154669970205"}],"defaultProperties":["branding_image","network_aware","Install Messages","Pending Install Messages","Pending Uninstall Messages","Uninstall Messages","Up-to-date Messages"],"options":{"remove_empty_properties":true},"properties":{"branding_image":{"default":null,"description":"Local path to image to use for branding, this will then be set either when Jamf Auto Update is installed or when Jamf Auto Update next runs.","options":{"infoText":"Key: branding_image","inputAttributes":{"placeholder":"/path/to/image"}},"title":"Branding Image","type":"string"},"network_aware":{"default":true,"description":"Default: true.","options":{"infoText":"Key: network_aware"},"title":"If true, Jamf Auto Update will not run when low data mode is enabled or when connected to a hotspot.","type":"boolean"},"Install Messages":{"defaultProperties":["notify_installed"],"description":"Settings for the install messages.","options":{"remove_empty_properties":true},"properties":{"notify_installed":{"default":true,"description":"Default: true.","options":{"infoText":"Key: notify_installed"},"title":"Notify the logged in user when a title has been installed.","type":"boolean"},"installed_notification_type":{"default":"banner","description":"Default: banner.","enum":["alert","banner"],"options":{"infoText":"Key: installed_notification_type"},"title":"Notification type for notifying a title has been installed.","type":"string"},"installed_message":{"default":"%DISPLAYNAME% %DISPLAYVERSION% has been installed.","description":"The placeholders %DISPLAYNAME% and %DISPLAYVERSION% can be used for these notifications. Default: %DISPLAYNAME% %DISPLAYVERSION% has been installed.","format":"textarea","options":{"infoText":"Key: notify_installed_message"},"title":"Message to display when a title has been installed.","type":"string"}},"property_order":2,"required":["notify_installed"],"type":"object"}}}`

	// Construct the ForcedSettings
	forcedSettings := jamfpro.ForcedSettings{
		Plist:         plistContent,
		JsonSchema:    jsonSchema,
		SchemaSource:  "prod-custom-setting-schemas",
		SchemaDomain:  "com.jamf.autoupdate",
		SchemaVersion: "1.3",
		SchemaVariant: "Jamf Auto Update.json",
	}

	// Construct the PayloadContentItem
	payloadItem := jamfpro.PayloadContentItem{
		PayloadType:      "com.apple.ManagedClient.preferences",
		Forced:           &forcedSettings,
		PreferenceDomain: "com.jamf.autoupdate",
	}

	// Create the ResourceConfigProfile
	profile := jamfpro.ResourceConfigProfile{
		PayloadContent: []jamfpro.PayloadContentItem{payloadItem},
		Level:          "SYSTEM",
	}

	// Call the CreateConfigProfile function
	result, err := client.CreateConfigProfileWithCustomSettingsSchema(&profile)
	if err != nil {
		log.Fatalf("Failed to create configuration profile: %v", err)
	}

	// Print the UUID of the created profile
	fmt.Printf("Successfully created macOS Configuration Profile with UUID: %s\n", result.UUID)

	// Optional: Get the newly created profile
	createdProfile, err := client.GetConfigProfileByPayloadUUID(result.UUID)
	if err != nil {
		log.Fatalf("Failed to get newly created profile: %v", err)
	}

	fmt.Printf("Retrieved profile: %+v\n", createdProfile)
}
