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

	// Example profile to be updated
	profileToUpdate := jamfpro.ResourceMobileDeviceEnrollmentProfile{
		General: jamfpro.MobileDeviceEnrollmentProfileSubsetGeneral{
			Name:        "Configurator Enrollment Profile",
			Description: "string",
		},
		Location: jamfpro.MobileDeviceEnrollmentProfileSubsetLocation{
			// Initialize with empty or specific values if required
			Username:     "",
			Realname:     "",
			RealName:     "",
			EmailAddress: "",
			Position:     "",
			Phone:        "",
			PhoneNumber:  "",
			Department:   "",
			Building:     "",
			Room:         0, // or specific room number
		},
		Purchasing: jamfpro.MobileDeviceEnrollmentProfileSubsetPurchasing{
			IsPurchased:          true,
			IsLeased:             false,
			PONumber:             "",
			Vendor:               "",
			ApplecareID:          "",
			PurchasePrice:        "",
			PurchasingAccount:    "",
			PODate:               "",
			PODateEpoch:          0,
			PODateUTC:            "",
			WarrantyExpires:      "",
			WarrantyExpiresEpoch: 0,
			WarrantyExpiresUTC:   "",
			LeaseExpires:         "",
			LeaseExpiresEpoch:    0,
			LeaseExpiresUTC:      "",
			LifeExpectancy:       0,
			PurchasingContact:    "",
		},
	}

	profileID := "1" // Replace 123 with the actual ID

	updatedProfile, err := client.UpdateMobileDeviceEnrollmentProfileByID(profileID, &profileToUpdate)
	if err != nil {
		log.Fatalf("Error updating profile: %v", err)
	}

	fmt.Println("Updated Profile:", updatedProfile)
}
