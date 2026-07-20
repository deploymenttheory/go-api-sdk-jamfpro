package jamfpro

import (
	"errors"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-http-client/response"
)

// Mirrors a prestage: locks on the resource, on embedded subsets, on an
// optional pointer subset (the computer-side account settings), and in a slice.
type vlSubset struct {
	Name        string
	VersionLock int
}

type vlOptional struct {
	Enabled     bool
	VersionLock int
}

type vlResource struct {
	DisplayName string
	VersionLock int
	Location    vlSubset
	Purchasing  vlSubset
	Account     *vlOptional
	Items       []vlSubset
}

func TestSyncAllVersionLocks(t *testing.T) {
	current := &vlResource{
		DisplayName: "server",
		VersionLock: 7,
		Location:    vlSubset{Name: "loc", VersionLock: 3},
		Purchasing:  vlSubset{Name: "pur", VersionLock: 4},
		Account:     &vlOptional{Enabled: true, VersionLock: 5},
		Items:       []vlSubset{{VersionLock: 11}, {VersionLock: 12}},
	}
	request := &vlResource{
		DisplayName: "client",
		VersionLock: 999,
		Location:    vlSubset{Name: "loc-new", VersionLock: 999},
		Purchasing:  vlSubset{Name: "pur-new", VersionLock: 999},
		Account:     &vlOptional{Enabled: false, VersionLock: 999},
		Items:       []vlSubset{{VersionLock: 999}, {VersionLock: 999}},
	}

	syncAllVersionLocks(current, request)

	checks := []struct {
		name string
		got  int
		want int
	}{
		{"top-level", request.VersionLock, 7},
		{"embedded subset", request.Location.VersionLock, 3},
		{"second embedded subset", request.Purchasing.VersionLock, 4},
		{"pointer subset", request.Account.VersionLock, 5},
		{"slice element 0", request.Items[0].VersionLock, 11},
		{"slice element 1", request.Items[1].VersionLock, 12},
	}
	for _, c := range checks {
		if c.got != c.want {
			t.Errorf("%s lock: got %d, want %d", c.name, c.got, c.want)
		}
	}

	// Non-lock fields must survive untouched.
	if request.DisplayName != "client" {
		t.Errorf("DisplayName overwritten: got %q", request.DisplayName)
	}
	if request.Location.Name != "loc-new" {
		t.Errorf("subset payload overwritten: got %q", request.Location.Name)
	}
	if request.Account.Enabled {
		t.Error("pointer subset payload overwritten")
	}
}

// A subset's lock is its own; it must never be derived from the parent's.
// Deriving it (e.g. parent+1) is what produces OPTIMISTIC_LOCK_FAILED.
func TestSyncAllVersionLocksUsesSubsetOwnValue(t *testing.T) {
	current := &vlResource{
		VersionLock: 5,
		Location:    vlSubset{VersionLock: 0},
		Purchasing:  vlSubset{VersionLock: 0},
	}
	request := &vlResource{}

	syncAllVersionLocks(current, request)

	if request.Location.VersionLock != 0 || request.Purchasing.VersionLock != 0 {
		t.Errorf("subset locks must come from the subset, not the parent: got location=%d purchasing=%d, want 0 and 0",
			request.Location.VersionLock, request.Purchasing.VersionLock)
	}
	if request.VersionLock != 5 {
		t.Errorf("top-level lock: got %d, want 5", request.VersionLock)
	}
}

func TestSyncAllVersionLocksEdgeCases(t *testing.T) {
	t.Run("nil pointer subsets on both sides", func(t *testing.T) {
		current := &vlResource{VersionLock: 2}
		request := &vlResource{}
		syncAllVersionLocks(current, request)
		if request.VersionLock != 2 {
			t.Errorf("got %d, want 2", request.VersionLock)
		}
	})

	t.Run("server has subset, request does not", func(t *testing.T) {
		current := &vlResource{VersionLock: 2, Account: &vlOptional{VersionLock: 9}}
		request := &vlResource{}
		syncAllVersionLocks(current, request)
		if request.Account != nil {
			t.Error("must not fabricate a subset the request omitted")
		}
	})

	t.Run("mismatched slice lengths", func(t *testing.T) {
		current := &vlResource{Items: []vlSubset{{VersionLock: 1}}}
		request := &vlResource{Items: []vlSubset{{VersionLock: 999}, {VersionLock: 888}}}
		syncAllVersionLocks(current, request)
		if request.Items[0].VersionLock != 1 {
			t.Errorf("overlapping element: got %d, want 1", request.Items[0].VersionLock)
		}
		if request.Items[1].VersionLock != 888 {
			t.Errorf("extra element must be untouched: got %d, want 888", request.Items[1].VersionLock)
		}
	})

	t.Run("mismatched types copy nothing", func(t *testing.T) {
		request := &vlResource{VersionLock: 42}
		syncAllVersionLocks(&vlSubset{VersionLock: 1}, request)
		if request.VersionLock != 42 {
			t.Errorf("cross-type copy occurred: got %d, want 42", request.VersionLock)
		}
	})
}

func TestZeroAllVersionLocks(t *testing.T) {
	request := &vlResource{
		DisplayName: "keep me",
		VersionLock: 9,
		Location:    vlSubset{VersionLock: 9},
		Purchasing:  vlSubset{VersionLock: 9},
		Account:     &vlOptional{VersionLock: 9},
		Items:       []vlSubset{{VersionLock: 9}},
	}

	zeroAllVersionLocks(request)

	got := []int{
		request.VersionLock,
		request.Location.VersionLock,
		request.Purchasing.VersionLock,
		request.Account.VersionLock,
		request.Items[0].VersionLock,
	}
	for i, v := range got {
		if v != newResourceVersionLock {
			t.Errorf("lock %d: got %d, want %d", i, v, newResourceVersionLock)
		}
	}
	if request.DisplayName != "keep me" {
		t.Errorf("payload altered: got %q", request.DisplayName)
	}
}

func TestTopVersionLock(t *testing.T) {
	if got, ok := topVersionLock(&vlResource{VersionLock: 6}); !ok || got != 6 {
		t.Errorf("got (%d, %v), want (6, true)", got, ok)
	}
	if _, ok := topVersionLock(nil); ok {
		t.Error("nil must report no lock")
	}
	var nilRes *vlResource
	if _, ok := topVersionLock(nilRes); ok {
		t.Error("typed nil must report no lock")
	}
	if _, ok := topVersionLock(&struct{ Name string }{}); ok {
		t.Error("struct without a lock must report none")
	}
}

func TestIsVersionLockConflict(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"409 conflict", &response.APIError{StatusCode: http.StatusConflict}, true},
		{"412 precondition failed", &response.APIError{StatusCode: http.StatusPreconditionFailed}, true},
		{"500 server error", &response.APIError{StatusCode: http.StatusInternalServerError}, false},
		{"400 bad request", &response.APIError{StatusCode: http.StatusBadRequest}, false},
		{"non-API error", errors.New("boom"), false},
		{"nil", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVersionLockConflict(tt.err); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRetryableVersionLockError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"500", &response.APIError{StatusCode: http.StatusInternalServerError}, true},
		{"503", &response.APIError{StatusCode: http.StatusServiceUnavailable}, true},
		{"404", &response.APIError{StatusCode: http.StatusNotFound}, false},
		{"401", &response.APIError{StatusCode: http.StatusUnauthorized}, false},
		{"non-API error", errors.New("boom"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRetryableVersionLockError(tt.err); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// The real prestage types must be walked correctly by the reflection, including
// the computer-side account settings that has no mobile counterpart.
func TestSyncAllVersionLocksOnRealPrestageTypes(t *testing.T) {
	t.Run("computer", func(t *testing.T) {
		current := &ResourceComputerPrestage{
			VersionLock:           4,
			LocationInformation:   ComputerPrestageSubsetLocationInformation{VersionLock: 1},
			PurchasingInformation: ComputerPrestageSubsetPurchasingInformation{VersionLock: 2},
			AccountSettings:       ComputerPrestageSubsetAccountSettings{VersionLock: 3},
		}
		request := &ResourceComputerPrestage{
			VersionLock:           99,
			LocationInformation:   ComputerPrestageSubsetLocationInformation{VersionLock: 99},
			PurchasingInformation: ComputerPrestageSubsetPurchasingInformation{VersionLock: 99},
			AccountSettings:       ComputerPrestageSubsetAccountSettings{VersionLock: 99},
		}

		syncAllVersionLocks(current, request)

		if request.VersionLock != 4 {
			t.Errorf("top: got %d, want 4", request.VersionLock)
		}
		if request.LocationInformation.VersionLock != 1 {
			t.Errorf("location: got %d, want 1", request.LocationInformation.VersionLock)
		}
		if request.PurchasingInformation.VersionLock != 2 {
			t.Errorf("purchasing: got %d, want 2", request.PurchasingInformation.VersionLock)
		}
		if request.AccountSettings.VersionLock != 3 {
			t.Errorf("account settings: got %d, want 3", request.AccountSettings.VersionLock)
		}
	})

	t.Run("mobile", func(t *testing.T) {
		current := &ResourceMobileDevicePrestage{
			VersionLock:           4,
			LocationInformation:   MobileDevicePrestageSubsetLocationInformation{VersionLock: 1},
			PurchasingInformation: MobileDevicePrestageSubsetPurchasingInformation{VersionLock: 2},
		}
		request := &ResourceMobileDevicePrestage{
			VersionLock:           99,
			LocationInformation:   MobileDevicePrestageSubsetLocationInformation{VersionLock: 99},
			PurchasingInformation: MobileDevicePrestageSubsetPurchasingInformation{VersionLock: 99},
		}

		syncAllVersionLocks(current, request)

		if request.VersionLock != 4 {
			t.Errorf("top: got %d, want 4", request.VersionLock)
		}
		if request.LocationInformation.VersionLock != 1 {
			t.Errorf("location: got %d, want 1", request.LocationInformation.VersionLock)
		}
		if request.PurchasingInformation.VersionLock != 2 {
			t.Errorf("purchasing: got %d, want 2", request.PurchasingInformation.VersionLock)
		}
	})
}
