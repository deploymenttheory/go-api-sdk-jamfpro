package jamfpro

import (
	"testing"

	"go.uber.org/zap"
)

func TestLoadConfigFromEnvReadsPlatformScopeIDs(t *testing.T) {
	t.Setenv("TENANT_ID", "tenant-id")
	t.Setenv("ENVIRONMENT_ID", "environment-id")

	config, err := loadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if config.TenantID != "tenant-id" {
		t.Errorf("TenantID = %q, want %q", config.TenantID, "tenant-id")
	}
	if config.EnvironmentID != "environment-id" {
		t.Errorf("EnvironmentID = %q, want %q", config.EnvironmentID, "environment-id")
	}
}

func TestInitializeAPIIntegrationRequiresExactlyOnePlatformScope(t *testing.T) {
	for _, tt := range []struct {
		name          string
		tenantID      string
		environmentID string
	}{
		{name: "missing"},
		{name: "both", tenantID: "tenant-id", environmentID: "environment-id"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := initializeAPIIntegration(&ConfigContainer{
				AuthMethod:    "platform",
				TenantID:      tt.tenantID,
				EnvironmentID: tt.environmentID,
			}, zap.NewNop().Sugar())
			if err == nil {
				t.Fatal("initializeAPIIntegration() succeeded with an ambiguous platform scope")
			}
		})
	}
}
