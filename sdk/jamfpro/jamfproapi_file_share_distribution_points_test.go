package jamfpro

import (
	"encoding/json"
	"reflect"
	"testing"
)

// This testing function could be done better, as marshalling in an out of maps can cause formatting differences
// which will not be representative of how the api would recieve the unmarshalled struct. 
// @TODO: make use of jsonfiles in the mocks folder to import and export configurations.
func TestResourceFileShareDistributionPoint_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    ResourceFileShareDistributionPoint
		expected map[string]any
	}{
		// {
		// 	name: "Full struct with all fields",
		// 	input: ResourceFileShareDistributionPoint{
		// 		ShareName:                 "TestShare",
		// 		Workgroup:                 "TestWorkgroup",
		// 		Port:                      445,
		// 		ReadWriteUsername:         "rwuser",
		// 		ReadWritePassword:         "rwpass",
		// 		ReadOnlyUsername:          "rouser",
		// 		ReadOnlyPassword:          "ropass",
		// 		ID:                        "1",
		// 		Name:                      "Test Distribution Point",
		// 		ServerName:                "server.example.com",
		// 		Principal:                 true,
		// 		BackupDistributionPointID: "2",
		// 		SSHUsername:               "sshuser",
		// 		SSHPassword:               "sshpass",
		// 		LocalPathToShare:          "/path/to/share",
		// 		FileSharingConnectionType: "SMB",
		// 		HTTPSEnabled:              true,
		// 		HTTPSPort:                 443,
		// 		HTTPSContext:              "context",
		// 		HTTPSSecurityType:         "BASIC",
		// 		HTTPSUsername:             "httpsuser",
		// 		HTTPSPassword:             "httpspass",
		// 		EnableLoadBalancing:       true,
		// 	},
		// 	expected: map[string]interface{}{
		// 		"shareName":                 "TestShare",
		// 		"workgroup":                 "TestWorkgroup",
		// 		"port":                      float64(445),
		// 		"readWriteUsername":         "rwuser",
		// 		"readWritePassword":         "rwpass",
		// 		"readOnlyUsername":          "rouser",
		// 		"readOnlyPassword":          "ropass",
		// 		"id":                        "1",
		// 		"name":                      "Test Distribution Point",
		// 		"serverName":                "server.example.com",
		// 		"principal":                 true,
		// 		"backupDistributionPointId": "2",
		// 		"sshUsername":               "sshuser",
		// 		"sshPassword":               "sshpass",
		// 		"localPathToShare":          "/path/to/share",
		// 		"fileSharingConnectionType": "SMB",
		// 		"httpsEnabled":              true,
		// 		"httpsPort":                 float64(443),
		// 		"httpsContext":              "context",
		// 		"httpsSecurityType":         "BASIC",
		// 		"httpsUsername":             "httpsuser",
		// 		"httpsPassword":             "httpspass",
		// 		"enableLoadBalancing":       true,
		// 	},
		// },
		{
			name: "Required keys are correctly zeroed and present",
			input: ResourceFileShareDistributionPoint{

			},
			expected: map[string]any{
				"fileSharingConnectionType":"",
				"httpsEnabled": false,
				"httpsPort": float64(0),
				"httpsSecurityType": "",
				"name": "",
				"serverName": "",
			},
		},
		{
			name: "Minimal required fields only",
			input: ResourceFileShareDistributionPoint{
				Name:                      "Minimal DP",
				ServerName:                "minimal.example.com",
				FileSharingConnectionType: "NONE",
				HTTPSEnabled:              false,
				HTTPSPort:                 443,
				HTTPSSecurityType:         "NONE",
			},
			// expected: "{\"name\":\"Minimal DP\",\"serverName\":\"minimal.example.com\",\"fileSharingConnectionType\":\"NONE\",\"httpsEnabled\":false,\"httpsPort\":443,\"httpsSecurityType\":\"NONE\"}",
			expected: map[string]any{
				"name":                      "Minimal DP",
				"serverName":                "minimal.example.com",
				"fileSharingConnectionType": "NONE",
				"httpsEnabled":              false,
				"httpsPort":                 float64(443),
				"httpsSecurityType":         "NONE",
			},
		},
		{
			name: "Maximal fields for SMB",
			input: ResourceFileShareDistributionPoint{
				Principal: true,
				BackupDistributionPointID: "3",
				FileSharingConnectionType: "SMB",
				Port: 139,
				HTTPSEnabled: true,
				HTTPSPort: 443,
				HTTPSSecurityType: "USERNAME_PASSWORD",
				EnableLoadBalancing: true,
				Name: "distributionpointname",
				ServerName: "servername",
				SSHUsername: "sshusername1",
				SSHPassword: "sshpassword1",
				LocalPathToShare: "/path/dir/file",
				ShareName: "sharename",
				Workgroup: "workgroupname",
				ReadWriteUsername: "readusername",
				ReadWritePassword: "readpassword1",
				ReadOnlyUsername: "readonlyusername",
				ReadOnlyPassword: "readonlypassword",
				HTTPSContext: "httpcontext",
				HTTPSUsername: "httpsusername",
				HTTPSPassword: "httpspassword",
			},
			expected: map[string]any{
				"principal": true,
				"backupDistributionPointId": "3",
				"fileSharingConnectionType": "SMB",
				"port": float64(139),
				"httpsEnabled": true,
				"httpsPort": float64(443),
				"httpsSecurityType": "USERNAME_PASSWORD",
				"enableLoadBalancing": true,
				"name": "distributionpointname",
				"serverName": "servername",
				"sshUsername": "sshusername1",
				"sshPassword": "sshpassword1",
				"localPathToShare": "/path/dir/file",
				"shareName": "sharename",
				"workgroup": "workgroupname",
				"readWriteUsername": "readusername",
				"readWritePassword": "readpassword1",
				"readOnlyUsername": "readonlyusername",
				"readOnlyPassword": "readonlypassword",
				"httpsContext": "httpcontext",
				"httpsUsername": "httpsusername",
				"httpsPassword": "httpspassword",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the input struct
			actualBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Errorf("Failed to marshal struct: %v", err)
				return
			}

			// Unmarshal actual result back to map for comparison
			var actualMap map[string]any
			if err := json.Unmarshal(actualBytes, &actualMap); err != nil {
				t.Errorf("Failed to unmarshal actual result: %v", err)
				return
			}

			// Compare the maps directly (order-independent)
			if !reflect.DeepEqual(actualMap, tt.expected) {
				actualJSON, _ := json.MarshalIndent(actualMap, "", "  ")
				expectedJSON, _ := json.MarshalIndent(tt.expected, "", "  ")
				t.Errorf("JSON marshalling mismatch:\nwant: %s\ngot:  %s", expectedJSON, actualJSON)
			}
		})
	}
}
