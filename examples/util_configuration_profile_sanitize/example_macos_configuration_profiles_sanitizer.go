package main

import (
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/utils"
)

func main() {
	target_filepath := "examples/util_configuration_profile_sanitize/payload.mobileconfig"

	configProfile, err := utils.ConfigurationFilePlistToStructFromFile(target_filepath)
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := json.MarshalIndent(configProfile, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

}
