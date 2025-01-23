package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

type ExportData struct {
	Version       string    `json:"version"`
	ExportDate    time.Time `json:"export_date"`
	JSSObjects    []string  `json:"jss_objects"`
	JSSSettings   []string  `json:"jss_settings"`
	JSSActions    []string  `json:"jss_actions"`
	Recon         []string  `json:"recon,omitempty"`
	CasperAdmin   []string  `json:"casper_admin,omitempty"`
	CasperRemote  []string  `json:"casper_remote,omitempty"`
	CasperImaging []string  `json:"casper_imaging,omitempty"`
}

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	outputPath := "jamf_pro_admin_account_privileges.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	versionInfo, err := client.GetJamfProVersion()
	if err != nil {
		log.Fatalf("Failed to get version: %v", err)
	}

	accounts, err := client.GetAccounts()
	if err != nil {
		log.Fatalf("Failed to get accounts: %v", err)
	}

	var adminAccount *jamfpro.ResourceAccount
	for _, user := range accounts.Users {
		account, err := client.GetAccountByID(fmt.Sprint(user.ID))
		if err != nil {
			continue
		}
		if account.PrivilegeSet == "Administrator" {
			adminAccount = account
			break
		}
	}

	if adminAccount == nil {
		log.Fatal("No administrator account found")
	}

	export := ExportData{
		Version:       *versionInfo.Version,
		ExportDate:    time.Now().UTC(),
		JSSObjects:    adminAccount.Privileges.JSSObjects,
		JSSSettings:   adminAccount.Privileges.JSSSettings,
		JSSActions:    adminAccount.Privileges.JSSActions,
		Recon:         adminAccount.Privileges.Recon,
		CasperAdmin:   adminAccount.Privileges.CasperAdmin,
		CasperRemote:  adminAccount.Privileges.CasperRemote,
		CasperImaging: adminAccount.Privileges.CasperImaging,
	}

	data, err := json.MarshalIndent(export, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Printf("Successfully exported privileges to %s", outputPath)
}
