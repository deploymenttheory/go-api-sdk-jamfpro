// This example demonstrates how to initialize the Jamf Pro client using configurations loaded from environment variables.
package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

/*
Ensure environment variables are set before running in this mode.
You will need to set the following environment variables:

export LOG_LEVEL="warning"
export HIDE_SENSITIVE_DATA="true"
export INSTANCE_DOMAIN="your-instance-domain"
export AUTH_METHOD="oauth2"
export CLIENT_ID="your-client-id"
export CLIENT_SECRET="your-client-secret"
export BASIC_AUTH_USERNAME="your-basic-auth-username"
export BASIC_AUTH_PASSWORD="your-basic-auth-password"
export JAMF_LOAD_BALANCER_LOCK="true"
export MAX_RETRY_ATTEMPTS="3"
export ENABLE_DYNAMIC_RATE_LIMITING="false"
export MAX_CONCURRENT_REQUESTS="1"
export TOKEN_REFRESH_BUFFER_PERIOD_SECONDS="300"
export TOTAL_RETRY_DURATION_SECONDS="60"
export CUSTOM_TIMEOUT_SECONDS="60"
export FOLLOW_REDIRECTS="true"
export MAX_REDIRECTS="5"
export ENABLE_CONCURRENCY_MANAGEMENT="true"
export CUSTOM_COOKIES=""
export MANDATORY_REQUEST_DELAY_MILLISECONDS="0"
export RETRY_ELIGIABLE_REQUESTS="true"
*/

func main() {
	// Initialize the Jamf Pro client using configurations loaded from environment variables
	client, err := jamfpro.BuildClientWithEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call the GetComputers method
	computers, err := client.GetComputers()
	if err != nil {
		log.Fatalf("Error fetching computers: %v", err)
	}

	// Print out the fetched computers
	fmt.Println("Fetched Computers:")
	for _, computer := range computers.Results {
		fmt.Printf("ID: %d, Name: %s\n", computer.ID, computer.Name)
	}
}
