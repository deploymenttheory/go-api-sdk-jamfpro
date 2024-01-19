package main

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/utils"
)

func main() {
	target_filepath := "examples/util_configuration_profile_sanitize/payload.mobileconfig"

	plist, err := utils.SanitiseMacOsConfigurationProfile(target_filepath)

	if err != nil {
		fmt.Printf("problem: %v", err)
	}

	fmt.Println(plist)
}
