package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Path to the Jamf Pro client configuration file
	configFilePath = "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Computer ID to check for certificates
	computerID = "14"

	// Certificate name to check for expiration
	// This is the Common Name of the certificate you want to find
	certificateName = "ACME Corp Root CA"
)

func main() {
	// Initialize the Jamf Pro client with the configuration file
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get the computer inventory details which includes certificate information
	computerInventory, err := client.GetComputerInventoryByID(computerID)
	if err != nil {
		log.Fatalf("Failed to get computer inventory: %v", err)
	}

	// Check if certificates are available
	if computerInventory.Certificates == nil || len(computerInventory.Certificates) == 0 {
		log.Fatalf("No certificates found for computer ID %s", computerID)
	}

	// Search for the specified certificate by name
	found := false
	for _, cert := range computerInventory.Certificates {
		// Check if the certificate common name matches what we're looking for
		// Use case-insensitive comparison to be more forgiving
		if strings.EqualFold(cert.CommonName, certificateName) {
			found = true

			// Parse the expiration date
			expirationDate, err := parseJamfDate(cert.ExpirationDate)
			if err != nil {
				log.Printf("Warning: Could not parse expiration date for certificate '%s': %v", cert.CommonName, err)
				continue
			}

			// Calculate days until expiration
			daysUntilExpiration := daysBetween(time.Now(), expirationDate)

			// Output the certificate information
			fmt.Printf("Certificate: %s\n", cert.CommonName)
			fmt.Printf("Serial Number: %s\n", cert.SerialNumber)
			fmt.Printf("Expiration Date: %s\n", cert.ExpirationDate)
			fmt.Printf("Days Until Expiration: %d\n", daysUntilExpiration)
			fmt.Printf("Status: %s\n", cert.CertificateStatus)
			fmt.Printf("Lifecycle Status: %s\n", cert.LifecycleStatus)
			fmt.Printf("SHA-1 Fingerprint: %s\n", cert.Sha1Fingerprint)

			// Warn if expiration is less than 30 days away
			if daysUntilExpiration < 30 {
				fmt.Printf("WARNING: Certificate expires in less than 30 days!\n")
			}
		}
	}

	if !found {
		fmt.Printf("Certificate with name '%s' not found on computer ID %s\n", certificateName, computerID)
	}
}

// parseJamfDate parses a date string from Jamf Pro API
func parseJamfDate(dateStr string) (time.Time, error) {
	// Jamf Pro API returns dates in multiple formats
	// Try RFC3339 format first (e.g., "2023-12-07T08:46:10.819Z")
	t, err := time.Parse(time.RFC3339, dateStr)
	if err == nil {
		return t, nil
	}

	// Try another common format (e.g., "2023-12-07")
	t, err = time.Parse("2006-01-02", dateStr)
	if err == nil {
		return t, nil
	}

	// Try yet another format that Jamf sometimes uses (e.g., "Dec 7, 2023, 8:46 AM")
	t, err = time.Parse("Jan 2, 2006, 3:04 PM", dateStr)
	if err == nil {
		return t, nil
	}

	// If all formats fail, return the error from the last attempt
	return time.Time{}, fmt.Errorf("could not parse date '%s': %w", dateStr, err)
}

// daysBetween calculates the number of days between two dates
func daysBetween(from, to time.Time) int {
	// Normalize both times to the beginning of the day in UTC
	fromNorm := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, time.UTC)
	toNorm := time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, time.UTC)

	// Calculate the difference in days
	return int(toNorm.Sub(fromNorm).Hours() / 24)
}
