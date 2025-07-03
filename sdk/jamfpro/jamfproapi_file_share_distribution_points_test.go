package jamfpro

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResourceFileShareDistributionPoint_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    ResourceFileShareDistributionPoint
		expected map[string]interface{}
	}{
		{
			name: "Full struct with all fields",
			input: ResourceFileShareDistributionPoint{
				ShareName:                 "TestShare",
				Workgroup:                 "TestWorkgroup",
				Port:                      445,
				ReadWriteUsername:         "rwuser",
				ReadWritePassword:         "rwpass",
				ReadOnlyUsername:          "rouser",
				ReadOnlyPassword:          "ropass",
				ID:                        "1",
				Name:                      "Test Distribution Point",
				ServerName:                "server.example.com",
				Principal:                 true,
				BackupDistributionPointID: "2",
				SSHUsername:               "sshuser",
				SSHPassword:               "sshpass",
				LocalPathToShare:          "/path/to/share",
				FileSharingConnectionType: "SMB",
				HTTPSEnabled:              true,
				HTTPSPort:                 443,
				HTTPSContext:              "context",
				HTTPSSecurityType:         "BASIC",
				HTTPSUsername:             "httpsuser",
				HTTPSPassword:             "httpspass",
				EnableLoadBalancing:       true,
			},
			expected: map[string]interface{}{
				"shareName":                 "TestShare",
				"workgroup":                 "TestWorkgroup",
				"port":                      float64(445),
				"readWriteUsername":         "rwuser",
				"readWritePassword":         "rwpass",
				"readOnlyUsername":          "rouser",
				"readOnlyPassword":          "ropass",
				"id":                        "1",
				"name":                      "Test Distribution Point",
				"serverName":                "server.example.com",
				"principal":                 true,
				"backupDistributionPointId": "2",
				"sshUsername":               "sshuser",
				"sshPassword":               "sshpass",
				"localPathToShare":          "/path/to/share",
				"fileSharingConnectionType": "SMB",
				"httpsEnabled":              true,
				"httpsPort":                 float64(443),
				"httpsContext":              "context",
				"httpsSecurityType":         "BASIC",
				"httpsUsername":             "httpsuser",
				"httpsPassword":             "httpspass",
				"enableLoadBalancing":       true,
			},
		},
		{
			name: "Minimal required fields only",
			input: ResourceFileShareDistributionPoint{
				Name:                      "Minimal DP",
				ServerName:                "minimal.example.com",
				FileSharingConnectionType: "AFP",
				HTTPSEnabled:              false,
				HTTPSPort:                 443,
				HTTPSSecurityType:         "NONE",
			},
			expected: map[string]interface{}{
				"name":                      "Minimal DP",
				"serverName":                "minimal.example.com",
				"fileSharingConnectionType": "AFP",
				"httpsEnabled":              false,
				"httpsPort":                 float64(443),
				"httpsSecurityType":         "NONE",
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

			// Create a map from the actual JSON for comparison
			var actualMap map[string]interface{}
			err = json.Unmarshal(actualBytes, &actualMap)
			if err != nil {
				t.Errorf("Failed to unmarshal JSON: %v", err)
				return
			}

			// Compare the maps
			if !reflect.DeepEqual(actualMap, tt.expected) {
				t.Errorf("JSON marshalling mismatch:\nwant: %v\ngot:  %v", tt.expected, actualMap)
			}
		})
	}
}
