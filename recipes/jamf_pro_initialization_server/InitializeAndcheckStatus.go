package main

import (
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	maxRetries        = 30
	initialWaitTime   = 5 * time.Second
	maxWaitTime       = 2 * time.Minute
	completedStepCode = "SERVER_INIT_COMPLETE"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Initialize database connection
	dbPassword := "your-secure-database-password"
	fmt.Println("Initializing database connection...")
	if err := client.InitializeDatabaseConnection(dbPassword); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Database initialization successful")

	// Initialize Jamf Pro system
	systemConfig := &jamfpro.ResourceSystemInitialize{
		ActivationCode:  "XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX",
		InstitutionName: "My Organization",
		EulaAccepted:    true,
		Username:        "admin",
		Password:        "your-secure-admin-password",
		Email:           "admin@example.com",
		JssUrl:          "https://example.jamfcloud.com",
	}

	fmt.Println("Initializing Jamf Pro system...")
	if err := client.InitializeJamfProServer(systemConfig); err != nil {
		log.Fatalf("Failed to initialize Jamf Pro system: %v", err)
	}
	fmt.Println("System initialization triggered successfully")

	// Monitor startup status with exponential backoff
	startTime := time.Now()
	currentWait := initialWaitTime
	retryCount := 0

	fmt.Println("\nMonitoring startup status...")
	for retryCount < maxRetries {
		elapsed := time.Since(startTime).Round(time.Second)
		status, err := client.GetStartupStatus()

		if err != nil {
			fmt.Printf("Warning: Failed to get status: %v\n", err)
		} else {
			// Print current status
			fmt.Printf("\nTime elapsed: %v\n", elapsed)
			fmt.Printf("Current step: %s (%s)\n", status.Step, status.StepCode)
			fmt.Printf("Progress: %d%%\n", status.Percentage)

			// Check for warnings or errors
			if status.Warning != "" {
				fmt.Printf("Warning: %s (%s)\n", status.Warning, status.WarningCode)
			}
			if status.Error != "" {
				fmt.Printf("Error: %s (%s)\n", status.Error, status.ErrorCode)
			}

			// Check if initialization is complete
			if status.StepCode == completedStepCode {
				fmt.Printf("\nInitialization completed successfully after %v\n", elapsed)
				return
			}
		}

		// Sleep with exponential backoff
		time.Sleep(currentWait)
		currentWait *= 2
		if currentWait > maxWaitTime {
			currentWait = maxWaitTime
		}
		retryCount++
	}

	log.Fatalf("Initialization monitoring timed out after %v", time.Since(startTime))
}
