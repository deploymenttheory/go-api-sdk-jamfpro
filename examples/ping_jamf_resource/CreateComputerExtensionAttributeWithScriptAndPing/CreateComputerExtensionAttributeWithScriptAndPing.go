package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func loadScriptFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Load the script from a file
	scriptPath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/computer_extensioin_attribute.sh"
	scriptContent, err := loadScriptFromFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to load script from file: %v", err)
	}

	// Define the new computer extension attribute
	attribute := &jamfpro.ResourceComputerExtensionAttribute{
		Name:        "Computer Extension Attribute Script Test",
		Description: "Computer Extension Attribute SCript Test",
		DataType:    "String", // String / Integer / Date (YYYY-MM-DD hh:mm:ss)
		InputType: jamfpro.ComputerExtensionAttributeSubsetInputType{
			Type:     "Script",
			Script:   scriptContent,
			Platform: "Mac", // Set this to the desired platform: "Mac" or "Windows".
		},
		InventoryDisplay: "General", // General / Hardware / Operating System / User and Location / Purchasing / Extension Attribute
		ReconDisplay:     "Extension Attributes",
	}

	// Call CreateComputerExtensionAttribute function
	createdAttribute, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		log.Fatalf("Error creating Computer Extension Attribute: %v", err)
	}

	// Pretty print the created attribute in XML
	createdAttributeXML, err := xml.MarshalIndent(createdAttribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Created Computer Extension Attribute:\n", string(createdAttributeXML))

	// New step: Ping the created resource
	if createdAttribute.ID != 0 {
		resourceID := strconv.Itoa(createdAttribute.ID) // Convert the resource ID to a string
		endpoint := "/JSSResource/computerextensionattributes"

		// Perform the ping operation
		resp, err := client.PingResource(endpoint, resourceID)
		if err != nil {
			log.Fatalf("Error pinging Computer Extension Attribute with ID %s: %v", resourceID, err)
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Ping to Computer Extension Attribute with ID %s was successful.\n", resourceID)
		} else {
			fmt.Printf("Ping to Computer Extension Attribute with ID %s returned status code %d.\n", resourceID, resp.StatusCode)
		}
	} else {
		log.Println("No valid ID for the created Computer Extension Attribute. Ping operation skipped.")
	}
}
