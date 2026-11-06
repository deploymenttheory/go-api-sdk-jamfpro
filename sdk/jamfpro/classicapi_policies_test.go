package jamfpro

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestPolicyReadBackOmitsAbsentRemediateKeyType(t *testing.T) {
	// Jamf omits remediate_key_type from GET even when it was supplied on create.
	var policy ResourcePolicy
	if err := xml.Unmarshal([]byte(`<policy><disk_encryption><action>apply</action><disk_encryption_configuration_id>1</disk_encryption_configuration_id><auth_restart>false</auth_restart></disk_encryption></policy>`), &policy); err != nil {
		t.Fatal(err)
	}
	encoded, err := xml.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "remediate_key_type") {
		t.Fatal("a policy rebuilt from GET emits an empty remediate_key_type, which Jamf rejects with 409")
	}
}

func TestPolicyPreservesExplicitRemediateKeyType(t *testing.T) {
	for _, value := range []string{"Individual", "Institutional", "Individual And Institutional"} {
		t.Run(value, func(t *testing.T) {
			encoded, err := xml.Marshal(PolicySubsetDiskEncryption{RemediateKeyType: value})
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(encoded), "<remediate_key_type>"+value+"</remediate_key_type>") {
				t.Fatalf("explicit key type was lost: %s", encoded)
			}
		})
	}
}
