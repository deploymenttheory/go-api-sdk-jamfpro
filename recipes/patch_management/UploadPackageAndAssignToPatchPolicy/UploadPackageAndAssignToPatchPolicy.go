package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"path/filepath"

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

	// Configuration variables
	patchSoftwareTitle := "Mozilla Firefox"                           // Name of the patch software title in Jamf
	version := "134.0"                                                // Version to be patched
	pkgFilePath := "/Users/dafyddwatkins/Downloads/Firefox 134.0.pkg" // Path to the package file
	policyName := fmt.Sprintf("%s - %s", patchSoftwareTitle, version)

	fmt.Println("-------------------------------------------------")
	fmt.Printf("Starting patch policy workflow for %s version %s\n", patchSoftwareTitle, version)

	// Step 0: Check if the package already exists
	fmt.Printf("Checking if package '%s' already exists...\n", filepath.Base(pkgFilePath))

	// Check if the package exists by name
	existingPackage, err := client.GetPackageByName(filepath.Base(pkgFilePath))
	if err != nil {
		log.Printf("Error checking for existing package: %v", err)
	} else if existingPackage != nil {
		// If the package already exists, skip upload and show its ID
		fmt.Printf("Package '%s' already exists with ID: %s\n", filepath.Base(pkgFilePath), existingPackage.ID)
		return // Exit early since the package already exists
	} else {
		// If the package doesn't exist, proceed to upload
		fmt.Printf("Package '%s' does not exist. Proceeding with upload...\n", filepath.Base(pkgFilePath))
	}

	// Step 1: Upload the package
	fmt.Printf("Uploading package: %s\n", filepath.Base(pkgFilePath))

	pkg := &jamfpro.ResourcePackage{
		PackageName:          filepath.Base(pkgFilePath),
		FileName:             filepath.Base(pkgFilePath),
		CategoryID:           "-1",
		Priority:             3,
		FillUserTemplate:     jamfpro.BoolPtr(false),
		SWU:                  jamfpro.BoolPtr(false),
		RebootRequired:       jamfpro.BoolPtr(false),
		OSInstall:            jamfpro.BoolPtr(false),
		SuppressUpdates:      jamfpro.BoolPtr(false),
		SuppressFromDock:     jamfpro.BoolPtr(false),
		SuppressEula:         jamfpro.BoolPtr(false),
		SuppressRegistration: jamfpro.BoolPtr(false),
	}

	// Debug: Upload the package and show response
	uploadResponse, err := client.DoPackageUpload(pkgFilePath, pkg)
	if err != nil {
		log.Fatalf("Failed to upload package: %v", err)
	}
	fmt.Printf("Package uploaded successfully with ID: %s\n", uploadResponse.ID)

	// Step 2: Get the Patch Software Title Configuration
	softwareTitle, err := client.GetPatchSoftwareTitleConfigurationByName(patchSoftwareTitle)
	if err != nil {
		log.Fatalf("Failed to get patch software title configuration: %v", err)
	}

	// Debug: Print full software title configuration details
	softwareTitleJson, _ := json.MarshalIndent(softwareTitle, "", "  ")
	fmt.Printf("Found patch software title configuration with ID: %s\n", softwareTitle.ID)
	fmt.Printf("Software Title Configuration: %s\n", string(softwareTitleJson))

	// Step 3: After uploading the package, associate it with the patch software title configuration
	fmt.Println("Associating uploaded package with patch software title configuration...")

	// Assuming the uploaded package has an ID (from the upload response):
	packageID := uploadResponse.ID

	// Create the request to associate the package with the patch software title
	patchSoftwareTitleConfiguration := &jamfpro.ResourcePatchSoftwareTitleConfiguration{
		CategoryID:         "-1", // Set the appropriate category ID
		SiteID:             "-1", // Set the appropriate site ID
		UiNotifications:    true, // UI Notifications enabled (match Python behavior)
		EmailNotifications: true, // Email Notifications enabled (match Python behavior)
		Packages: []jamfpro.PatchSoftwareTitleConfigurationSubsetPackage{
			{
				PackageId:   packageID,                  // Package ID from uploaded package
				Version:     version,                    // Version of the package
				DisplayName: filepath.Base(pkgFilePath), // Package display name
			},
		},
	}

	// Perform a check to see if the patch software title configuration exists
	var response interface{}
	if softwareTitle.ID != "" {
		fmt.Println("Updating existing patch software title configuration...")
		response, err = client.UpdatePatchSoftwareTitleConfigurationById(softwareTitle.ID, *patchSoftwareTitleConfiguration)
		if err != nil {
			log.Fatalf("Failed to update patch software title configuration: %v", err)
		}
		fmt.Printf("Patch Software Title Configuration successfully updated with ID: %s\n", softwareTitle.ID)
	} else {
		fmt.Println("Creating new patch software title configuration...")
		responseCreate, err := client.CreatePatchSoftwareTitleConfiguration(*patchSoftwareTitleConfiguration)
		if err != nil {
			log.Fatalf("Failed to create patch software title configuration: %v", err)
		}
		fmt.Printf("Patch Software Title Configuration successfully created with ID: %s\n", responseCreate.ID)
		response = responseCreate
	}

	// Debug: Print the response for inspection
	responseJson, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("Patch Software Title Configuration Response: %s\n", string(responseJson))

	// Step 4: Get the Patch Software Title Definitions to retrieve missing values
	// For more information on how to add parameters to this request, see docs/url_queries.md
	softwareTitleDefinitions, err := client.GetPatchSoftwareTitleDefinitions(softwareTitle.ID, url.Values{})
	if err != nil {
		log.Fatalf("Failed to get patch software title definitions: %v", err)
	}

	// Step 5: Match the version with the provided version
	var matchedDefinition *jamfpro.ResourcePatchSoftwareTitleDefinition
	for _, def := range softwareTitleDefinitions.Results {
		if def.Version == version {
			matchedDefinition = &def
			break
		}
	}

	if matchedDefinition == nil {
		log.Fatalf("No matching patch software title definition found for version '%s'", version)
	}

	// Step 6: Create the patch policy request with matched values
	policyRequest := &jamfpro.ResourcePatchPolicyClassicAPI{
		General: jamfpro.ResourcePatchPolicyCreateRequestGeneral{
			Name:               policyName,
			Enabled:            true,
			TargetVersion:      version,
			ReleaseDate:        matchedDefinition.ReleaseDate, // Populate ReleaseDate
			IncrementalUpdates: false,
			Reboot:             matchedDefinition.RebootRequired,         // Populate RebootRequired
			MinimumOS:          matchedDefinition.MinimumOperatingSystem, // Populate MinimumOperatingSystem
			KillApps: jamfpro.ResourcePatchPolicyCreateRequestKillApps{
				KillApp: []jamfpro.ResourcePatchPolicyCreateRequestKillApp{
					{
						KillAppName:     "Firefox.app",
						KillAppBundleID: "org.mozilla.firefox",
					},
				},
			},
			DistributionMethod: "selfservice",
			AllowDowngrade:     false,
			PatchUnknown:       true,
		},
		Scope: jamfpro.ResourcePatchPolicyCreateRequestScope{
			AllComputers: true,
		},
		UserInteraction: jamfpro.ResourcePatchPolicyCreateRequestUserInteraction{
			InstallButtonText:      "Update",
			SelfServiceDescription: "Update Firefox to the latest version",
			Notifications: jamfpro.ResourcePatchPolicyCreateRequestNotifications{
				Enabled: true,
				Type:    "Self Service",
				Subject: "Firefox Update Available",
				Message: "A new version of Firefox is available in Self Service",
				Reminders: jamfpro.ResourcePatchPolicyCreateRequestReminders{
					Enabled:   true,
					Frequency: 1,
				},
			},
			Deadlines: jamfpro.ResourcePatchPolicyCreateRequestDeadlines{
				Enabled: true,
				Period:  7,
			},
			GracePeriod: jamfpro.ResourcePatchPolicyCreateRequestGracePeriod{
				Duration:            15,
				NotificationSubject: "Important",
				Message:             "$APP_NAMES will quit in $DELAY_MINUTES minutes so that $SOFTWARE_TITLE can be updated.",
			},
		},
		SoftwareTitleConfigurationID: softwareTitle.ID,
	}

	// Debug: Print the patch policy request details before creation
	policyRequestXml, _ := xml.MarshalIndent(policyRequest, "", "  ")
	fmt.Printf("Patch Policy Request: %s\n", string(policyRequestXml))

	// Create the patch policy
	err = client.CreatePatchPolicy(softwareTitle.ID, policyRequest)
	if err != nil {
		log.Fatalf("Failed to create patch policy: %v", err)
	}

	fmt.Printf("Successfully linked package '%s' with version '%s'\n", pkg.PackageName, version)

	// Get and display the created policy
	createdPolicy, err := client.GetPatchPolicyByName(policyName)
	if err != nil {
		log.Fatalf("Failed to get created patch policy: %v", err)
	}

	// Debug: Pretty print the created policy for detailed inspection
	createdPolicyJson, _ := json.MarshalIndent(createdPolicy, "", "  ")
	fmt.Printf("\nCreated Patch Policy Details:\n%s\n", string(createdPolicyJson))
	fmt.Println("-------------------------------------------------")
}
