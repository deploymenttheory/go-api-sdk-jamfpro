package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

/*
Ensure environment variables are set before running in this mode.
You will need to set the following environment variables:

export CLIENT_ID="your-client-id"
export CLIENT_SECRET="your-client-secret"
export INSTANCE_NAME="your-instance-name"
export OVERRIDE_BASE_DOMAIN=""
export API_TYPE="jamfpro"
export LOG_LEVEL="LogLevelDebug"
export LOG_OUTPUT_FORMAT="console"
export LOG_CONSOLE_SEPARATOR=" "
export HIDE_SENSITIVE_DATA="true"
export MAX_RETRY_ATTEMPTS="3"
export ENABLE_DYNAMIC_RATE_LIMITING="true"
export MAX_CONCURRENT_REQUESTS="5"
export TOKEN_REFRESH_BUFFER_PERIOD="5m"
export TOTAL_RETRY_DURATION="5m"
export CUSTOM_TIMEOUT="10s"
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
