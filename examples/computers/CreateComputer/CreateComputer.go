package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a new computer configuration
	newComputer := jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			Name:         "APIGWFYFGH",
			SerialNumber: "APIGWFYFGH",                           // Must be Unique
			UDID:         "APIFF74D-C6B7-5589-93A9-19E8BDFEDFGH", // Must be Unique
			RemoteManagement: jamfpro.ComputerSubsetGeneralRemoteManagement{
				Managed: true,
			},
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Location: jamfpro.ComputerSubsetLocation{
			// Populate location fields if necessary
		},
		Purchasing: jamfpro.ComputerSubsetPurchasing{
			IsPurchased:          true,
			IsLeased:             false,
			PoNumber:             "PO123ABC",
			Vendor:               "Computer Supplies Co.",
			ApplecareID:          "AC123456789",
			PurchasePrice:        "2000.00",
			PurchasingAccount:    "IT Budget",
			PoDate:               "2023-01-15",
			PoDateEpoch:          1673760000,
			PoDateUtc:            "2023-01-01T00:00:00.000+0000",
			WarrantyExpires:      "2026-01-15",
			WarrantyExpiresEpoch: 1773984000,
			WarrantyExpiresUtc:   "2030-01-01T00:00:00.000+0000",
			LeaseExpires:         "2030-01-01",
			LeaseExpiresEpoch:    0,
			LeaseExpiresUtc:      "2030-01-01T00:00:00.000+0000",
			LifeExpectancy:       4,
			PurchasingContact:    "Jane Smith",
		},
		ExtensionAttributes: []jamfpro.ComputerSubsetExtensionAttributes{
			{
				ID:    2,
				Value: "", // Set value if necessary
			},
		},
		Hardware: jamfpro.ComputerSubsetHardware{
			Make:                        "Apple",
			Model:                       "MacBook Pro",
			ModelIdentifier:             "MacBookPro11,4",
			OsName:                      "macOS",
			OsVersion:                   "10.15.7",
			OsBuild:                     "19H2",
			MasterPasswordSet:           false,
			ActiveDirectoryStatus:       "Not Connected",
			ServicePack:                 "",
			ProcessorType:               "Intel Core i7",
			ProcessorArchitecture:       "x64",
			ProcessorSpeed:              2200,
			ProcessorSpeedMhz:           2200,
			NumberProcessors:            1,
			NumberCores:                 4,
			TotalRam:                    8192,
			TotalRamMb:                  8192,
			BootRom:                     "220.0.0.0.0",
			BusSpeed:                    100,
			BusSpeedMhz:                 100,
			BatteryCapacity:             85,
			CacheSize:                   6,
			CacheSizeKb:                 6144,
			AvailableRamSlots:           0,
			OpticalDrive:                "None",
			NicSpeed:                    "1Gbit",
			SmcVersion:                  "2.41f2",
			BleCapable:                  true,
			SipStatus:                   "Enabled",
			GatekeeperStatus:            "Enabled",
			XprotectVersion:             "2099",
			InstitutionalRecoveryKey:    "ExampleKey",
			DiskEncryptionConfiguration: "APFS Encrypted",
			SoftwareUpdateDeviceID:      "ExampleDeviceID",
			IsAppleSilicon:              false,
			SupportsIosAppInstalls:      false,
			Filevault2Users:             []jamfpro.ComputerSubsetHardwareFileVault2Users{{User: "testuser"}},
			Storage: []jamfpro.ComputerSubsetHardwareStorage{
				{
					Disk:            "disk0",
					Model:           "APPLE SSD AP0256",
					Revision:        "11.0.1",
					SerialNumber:    "S123456789",
					Size:            256,
					DriveCapacityMb: 256000,
					ConnectionType:  "PCI",
					SmartStatus:     "Verified",
					Partitions: []jamfpro.ComputerSubsetHardwareStoragePartitions{
						{
							Name:                "Macintosh HD",
							Size:                256,
							Type:                "APFS",
							PartitionCapacityMb: 256000,
							PercentageFull:      60,
							//FilevaultStatus:      "Enabled",
							//FilevaultPercent:     100,
							//Filevault2Status:     "Enabled",
							Filevault2Percent:    100,
							BootDriveAvailableMb: 102400,
							LvgUUID:              "ExampleLvgUUID",
							LvUUID:               "ExampleLvUUID",
							PvUUID:               "ExamplePvUUID",
						},
					},
				},
			},
			MappedPrinters: []jamfpro.ComputerSubsetHardwareMappedPrinters{
				{
					Name:     "Office Printer",
					URI:      "lpd://192.168.1.100",
					Type:     "Laser",
					Location: "Office",
				},
			},
		},
	}

	// Call CreateComputer function
	createdComputer, err := client.CreateComputer(newComputer)
	if err != nil {
		log.Fatalf("Error creating computer: %v", err)
	}

	// Pretty print the created department in JSON
	createdComputerJSON, err := xml.MarshalIndent(createdComputer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created computer data: %v", err)
	}
	fmt.Println("Created Computer:\n", string(createdComputerJSON))
}
