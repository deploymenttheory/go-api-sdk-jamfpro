package jamfpro

import (
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	apiresponse "github.com/deploymenttheory/go-api-http-client/response"
)

func TestSmartComputerGroupV2ContextTypedErrors(t *testing.T) {
	client := newGroupTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"httpStatus":404,"errors":[{"code":"INVALID_ID","description":"Not Found"}]}`))
	})
	_, err := client.GetSmartComputerGroupByIDV2WithContext(context.Background(), "42")
	var apiErr *apiresponse.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		t.Fatalf("404 must retain its typed status for state removal: %v", err)
	}
}

func TestSmartComputerGroupV2ContextCancellation(t *testing.T) {
	release := make(chan struct{})
	client := newGroupTestClient(t, func(_ http.ResponseWriter, req *http.Request) {
		_, _ = io.Copy(io.Discard, req.Body)
		select {
		case <-req.Context().Done():
		case <-release:
		}
	})
	t.Cleanup(func() { close(release) })
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := client.CreateSmartComputerGroupV2WithContext(ctx, ResourceSmartComputerGroupV2{Name: "test"})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("create must honor the caller's operation deadline: %v", err)
	}
}

func TestSmartComputerGroupV2ContextCRUD(t *testing.T) {
	for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			calls := 0
			client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
				calls++
				want := "/api/v2/computer-groups/smart-groups"
				if method != http.MethodPost {
					want += "/42"
				}
				if req.Method != method || req.URL.Path != want {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				if method == http.MethodDelete {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				_, _ = w.Write([]byte(`{"id":"42","name":"test"}`))
			})
			var err error
			switch method {
			case http.MethodPost:
				_, err = client.CreateSmartComputerGroupV2WithContext(context.Background(), ResourceSmartComputerGroupV2{Name: "test"})
			case http.MethodPut:
				_, err = client.UpdateSmartComputerGroupByIDV2WithContext(context.Background(), "42", ResourceSmartComputerGroupV2{Name: "test"})
			case http.MethodDelete:
				err = client.DeleteSmartComputerGroupByIDV2WithContext(context.Background(), "42")
			}
			if err != nil || calls != 1 {
				t.Fatalf("expected one successful write, got calls=%d, error=%v", calls, err)
			}
		})
	}
}

func TestSmartComputerGroupV2ContextNamePagination(t *testing.T) {
	calls := 0
	client := newGroupTestClient(t, func(w http.ResponseWriter, req *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		if req.URL.Query().Get("page") == "0" {
			_, _ = w.Write([]byte(`{"totalCount":101,"results":[{"id":"1","name":"other"}]}`))
		} else {
			_, _ = w.Write([]byte(`{"totalCount":101,"results":[{"id":"42","name":"Target"}]}`))
		}
	})
	group, err := client.GetSmartComputerGroupByNameV2WithContext(context.Background(), "target")
	if err != nil || group.ID != "42" || calls != 2 {
		t.Fatalf("lookup must page and compare exact names case-insensitively: group=%v, err=%v, calls=%d", group, err, calls)
	}
}
