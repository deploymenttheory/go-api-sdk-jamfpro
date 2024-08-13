package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"go.uber.org/zap"
)

func main() {
	// Define the path to the JSON configuration file
	logger, _ := zap.NewDevelopment()
	clientIntegration, _ := jamfprointegration.BuildWithOAuth(
		"https://lbgsandbox.jamfcloud.com",
		logger.Sugar(),
		5*time.Second,
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		false,
		&httpclient.ProdExecutor{Client: &http.Client{}},
	)
	clientConfig := httpclient.ClientConfig{
		Integration:           clientIntegration,
		Sugar:                 logger.Sugar(),
		PopulateDefaultValues: true,
		HideSensitiveData:     false,
		MaxRetryAttempts:      5,
		HTTPExecutor:          &httpclient.ProdExecutor{Client: &http.Client{}},
	}

	libClient, err := clientConfig.Build()

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	client := jamfpro.Client{HTTP: libClient}

	// Path to the icon file you want to upload
	filePath := "/Users/joseph/github/go-api-sdk-jamfpro/examples/icon/UploadIcon/cat.png"

	// Call the UploadIcon function
	uploadResponse, err := client.UploadIcon(filePath)
	if err != nil {
		fmt.Printf("Error uploading icon: %s\n", err)
		return
	}

	log.Println(uploadResponse)
}
