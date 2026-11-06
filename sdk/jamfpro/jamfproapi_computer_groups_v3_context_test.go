package jamfpro

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	apiresponse "github.com/deploymenttheory/go-api-http-client/response"
	"go.uber.org/zap"
)

type groupTestIntegration struct{ baseURL string }

func (i *groupTestIntegration) GetFQDN() string                     { return i.baseURL }
func (i *groupTestIntegration) ConstructURL(endpoint string) string { return i.baseURL + endpoint }
func (i *groupTestIntegration) GetAuthMethodDescriptor() string     { return "test" }
func (i *groupTestIntegration) CheckRefreshToken() error            { return nil }
func (i *groupTestIntegration) PrepRequestParamsAndAuth(req *http.Request) error {
	if len(req.URL.Path) >= 12 && req.URL.Path[:12] == "/JSSResource" {
		req.Header.Set("Content-Type", "application/xml")
	} else {
		req.Header.Set("Content-Type", "application/json")
	}
	return nil
}
func (i *groupTestIntegration) GetSessionCookies() ([]*http.Cookie, error) { return nil, nil }
func (i *groupTestIntegration) PrepRequestBody(body any, _, endpoint string) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	if len(endpoint) >= 12 && endpoint[:12] == "/JSSResource" {
		return xml.Marshal(body)
	}
	return json.Marshal(body)
}
func (i *groupTestIntegration) MarshalMultipartRequest(map[string]string, map[string]string) ([]byte, string, error) {
	return nil, "", nil
}

func newGroupTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	config := &httpclient.ClientConfig{
		Integration: &groupTestIntegration{baseURL: server.URL}, Sugar: zap.NewNop().Sugar(),
		Timeout: time.Second, HTTP: http.Client{},
	}
	transport, err := config.Build()
	if err != nil {
		t.Fatal(err)
	}
	return &Client{HTTP: transport}
}

func TestUpdateGroupDescriptionExplicitClear(t *testing.T) {
	client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch || req.URL.Path != "/api/v2/groups/platform-uuid" {
			t.Errorf("unexpected request %s %s", req.Method, req.URL.Path)
		}
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		if len(body) != 1 || body["groupDescription"] != "" {
			t.Errorf("description-only clear must retain explicit empty string, got %#v", body)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	if err := client.UpdateGroupDescriptionByIDV2WithContext(context.Background(), "platform-uuid", ""); err != nil {
		t.Fatal(err)
	}
}

func TestSmartComputerGroupPreservesTypedAPIError(t *testing.T) {
	client := newGroupTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"errors":[{"code":"INVALID_FIELD","description":"The criterion Test is not valid"}]}`))
	})
	_, err := client.CreateSmartComputerGroupV3WithContext(context.Background(), ResourceSmartComputerGroupV3{Name: "test"})
	var apiErr *apiresponse.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 400 || len(apiErr.Errors) != 1 {
		t.Fatalf("expected wrapped structured API error, got %v", err)
	}
}

func TestSmartComputerGroupNameLookupPaginates(t *testing.T) {
	pages := 0
	client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
		pages++
		w.Header().Set("Content-Type", "application/json")
		if req.URL.Query().Get("page") == "0" {
			_, _ = w.Write([]byte(`{"totalCount":101,"results":[{"id":"1","name":"other"}]}`))
			return
		}
		_, _ = w.Write([]byte(`{"totalCount":101,"results":[{"id":"42","name":"Example"}]}`))
	})
	group, err := client.GetSmartComputerGroupByNameV3WithContext(context.Background(), "example")
	if err != nil || group == nil || group.ID != "42" || pages != 2 {
		t.Fatalf("expected match on second page, got group=%v error=%v pages=%d", group, err, pages)
	}
}

func TestComputerGroupPlatformLookupUsesComputerPoolAndBothIDs(t *testing.T) {
	client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Query().Get("filter") != `groupType=="COMPUTER" and isSmart=="true"` {
			t.Errorf("unexpected group pool filter: %s", req.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"totalCount":1,"results":[{"groupType":"COMPUTER","groupJamfProId":"42","groupPlatformId":"platform-uuid"}]}`))
	})
	id, err := client.GetComputerSmartGroupPlatformIDByJamfProIDWithContext(context.Background(), "42")
	if err != nil || id != "platform-uuid" {
		t.Fatalf("got %q, %v", id, err)
	}
}

func TestClassicComputerGroupSingleWrite(t *testing.T) {
	calls := 0
	client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
		calls++
		var input struct {
			Name     string                 `xml:"name"`
			Site     SharedResourceSite     `xml:"site"`
			Criteria []SharedSubsetCriteria `xml:"criteria>criterion"`
		}
		if err := xml.NewDecoder(req.Body).Decode(&input); err != nil {
			t.Error(err)
		}
		if input.Name != "complete" || input.Site.ID != 3 || len(input.Criteria) != 2 {
			t.Errorf("partial group definition: %#v", input)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":42}`))
	})
	criteria := []SharedSubsetCriteria{{Name: "one", Priority: 0}, {Name: "two", Priority: 1}}
	_, err := client.UpdateComputerGroupByIDWithContext(context.Background(), "42", &ResourceComputerGroup{
		Name: "complete", IsSmart: true, Site: &SharedResourceSite{ID: 3},
		Criteria: &ComputerGroupSubsetContainerCriteria{Size: 2, Criterion: &criteria},
	})
	if err != nil || calls != 1 {
		t.Fatalf("expected one write, got %d, %v", calls, err)
	}
}
