package modules

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Function to generate a 26-digit random Recovery Lock password
func GenerateRandomRecoveryLockPassword() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// First digit between 0-9
	firstDigit := r.Intn(10)

	// 5 groups of 5 random digits (0-65535)
	groups := make([]int, 5)
	for i := range groups {
		groups[i] = r.Intn(65536) % 100000 // Ensure it's 5 digits by modulo 100000
	}

	// Format the password as a 26-digit string
	return fmt.Sprintf("%d%05d%05d%05d%05d%05d",
		firstDigit, groups[0], groups[1], groups[2], groups[3], groups[4])
}

// GetManagementIDByDeviceName retrieves the management ID for a device by its name.
// For more information on how to add parameters to this request, see docs/url_queries.md
func GetManagementIDByDeviceName(client *jamfpro.Client, deviceName string) (string, error) {
	inventories, err := client.GetComputersInventory(url.Values{})
	if err != nil {
		return "", fmt.Errorf("failed to get computer inventory: %w", err)
	}

	for _, inventory := range inventories.Results {
		if inventory.General.Name == deviceName {
			return inventory.General.ManagementId, nil
		}
	}

	return "", fmt.Errorf("device with name %s not found", deviceName)
}
