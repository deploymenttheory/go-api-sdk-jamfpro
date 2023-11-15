// jamfproapi_jamf_pro_sso_failover_test.go
// Jamf Pro Api - Jamf Pro SSO Failover Integration Testing
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover

/*
	Test Strategy:

Global Setup: The testing process begins with the initialization of the Jamf Pro HTTP Client.
This global setup phase involves creating a temporary API test role and setting up a corresponding
API client for integration testing. This ensures that all tests run in a consistent
and controlled environment.

Individual Test Execution: Each integration test, managed by testing.T, is executed according
to a predefined test plan. These tests utilize the temporary API client and test role established
in the setup phase. The use of testing.T facilitates granular error reporting and isolated
testing of specific functionalities within the Jamf Pro integration.

Global Teardown: Upon completion of all tests, the suite enters the teardown phase. This
involves a systematic cleanup of all test-generated resources, including the removal of
the temporary API client and the test role. This step is crucial for ensuring that the
testing environment is reset and no residual data impacts subsequent test runs.
*/
package jamfpro_integration_test

/*

// TestJamfProIntegration_GetSSOFailoverSettings tests the GetSSOFailoverSettings functionality of the Jamf Pro intTestClient.
// It verifies that the SSO failover settings can be retrieved correctly and asserts that the
// returned settings contain expected data. This test validates the ability
// of the client to interact with the Jamf Pro API and retrieve SSO failover information.
func TestJamfProIntegration_GetSSOFailoverSettings(t *testing.T) {

	failoverSettings, err := intTestClient.GetSSOFailoverSettings()
	if err != nil {
		log.Fatalf("Failed to get SSO failover settings: %v", err)
	}

	// Assert that failover URL is not nil and not empty
	if failoverSettings.FailoverURL == "" {
		t.Errorf("Expected a failover URL, got an empty string")
	}

	// Assert that generation time is not zero (assuming it's a Unix timestamp or similar)
	if failoverSettings.GenerationTime == 0 {
		t.Errorf("Expected a non-zero generation time, got zero")
	}

	// Log the retrieved failover settings for verification
	log.Printf("Retrieved SSO Failover URL: %s", failoverSettings.FailoverURL)
	log.Printf("Retrieved Generation Time: %d", failoverSettings.GenerationTime)
}

// TestJamfProIntegration_UpdateFailoverUrl tests the UpdateFailoverUrl functionality of the Jamf Pro intTestClient.
// It verifies that the SSO failover URL can be updated correctly and asserts that the
// returned settings contain the new failover URL and a new generation time.
func TestJamfProIntegration_UpdateFailoverUrl(t *testing.T) {

	// Update the SSO failover URL
	updatedFailoverSettings, err := intTestClient.UpdateFailoverUrl()
	if err != nil {
		log.Fatalf("Error updating SSO failover URL: %v", err)
	}

	// Assert that the updated failover URL is not empty
	if updatedFailoverSettings.FailoverURL == "" {
		t.Errorf("Expected a non-empty failover URL, got an empty string")
	}

	// Assert that the generation time is updated (not zero)
	if updatedFailoverSettings.GenerationTime == 0 {
		t.Errorf("Expected a non-zero generation time, got zero")
	}

	// Log the updated failover settings for verification
	log.Printf("Updated SSO Failover URL: %s", updatedFailoverSettings.FailoverURL)
	log.Printf("New Generation Time: %d", updatedFailoverSettings.GenerationTime)
}

*/
