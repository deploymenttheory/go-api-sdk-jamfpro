package jamfpro

import (
	"context"
	"errors"
	"net/http"
	"testing"

	apiresponse "github.com/deploymenttheory/go-api-http-client/response"
)

func TestSmartMobileDeviceGroupContextTypedErrors(t *testing.T) {
	for _, method := range []string{http.MethodGet, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
				if req.Method != method || req.URL.Path != "/api/v1/mobile-device-groups/smart-groups/42" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"httpStatus":404,"errors":[{"code":"INVALID_ID","description":"Not Found"}]}`))
			})
			var err error
			if method == http.MethodGet {
				_, err = client.GetSmartMobileDeviceGroupByIDV1WithContext(context.Background(), "42")
			} else {
				err = client.DeleteSmartMobileDeviceGroupByIDV1WithContext(context.Background(), "42")
			}
			var apiErr *apiresponse.APIError
			if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
				t.Fatalf("expected wrapped typed 404: %v", err)
			}
		})
	}
}
