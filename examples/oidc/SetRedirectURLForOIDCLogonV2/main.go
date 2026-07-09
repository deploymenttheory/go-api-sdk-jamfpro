package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// v2 OIDC login dispatch can return multiple IdP redirect options.
	request := &jamfpro.ResourceOIDCRedirectURL{
		OriginalURL:  "aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA==",
		EmailAddress: "admin@domain.name",
	}

	response, err := client.SetRedirectURLForOIDCLogonV2(request)
	if err != nil {
		log.Fatalf("Error dispatching OIDC v2 login: %v", err)
	}

	out, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling OIDC v2 dispatch response: %v", err)
	}
	fmt.Println("OIDC v2 Login Dispatch:\n", string(out))
}
