// shared_version_lock.go
// Support for Jamf Pro API endpoints that use optimistic locking.
//
// Certain resources — computer prestages, mobile device prestages, and their
// scope sub-resources — guard against concurrent modification with a
// versionLock token. The protocol is:
//
//  1. GET the resource. The response carries the current versionLock.
//  2. Echo that exact value back in the body of the write.
//  3. Creates (POST) must send 0, having no prior state to echo.
//
// If the value supplied is not the server's current one, the write is rejected
// with 409 Conflict / OPTIMISTIC_LOCK_FAILED.
//
// Jamf Pro docs: https://developer.jamf.com/jamf-pro/docs/optimistic-locking
package jamfpro

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/deploymenttheory/go-api-http-client/response"
)

const (
	// newResourceVersionLock is the versionLock a create must carry.
	newResourceVersionLock = 0

	// versionLockFieldName is the Go field name holding a lock token. Jamf
	// marshals it as "versionLock".
	versionLockFieldName = "VersionLock"

	// defaultVersionLockAttempts bounds how many times a rejected write is
	// re-driven. Each attempt costs one GET and one write.
	defaultVersionLockAttempts = 3
)

// syncAllVersionLocks copies every versionLock in current onto the matching
// field in request, walking the whole struct tree.
//
// These resources carry locks at more than one level: the resource itself plus
// each nested subset (location information, purchasing information, account
// settings). Every one must echo the value from the most recent GET or the
// write is rejected, and each subset holds its own value — a subset's lock is
// unrelated to its parent's and must never be derived from it.
//
// Listing those fields by hand at each call site is what lets a newly added
// subset ship with an unsynchronised lock, so they are discovered reflectively.
// Fields are matched by name and walked in parallel; a field present in one
// struct but not the other is skipped.
//
// current and request must be pointers to the same struct type; anything else
// is a no-op.
func syncAllVersionLocks(current, request any) {
	c, r := reflect.ValueOf(current), reflect.ValueOf(request)
	if !c.IsValid() || !r.IsValid() || c.Type() != r.Type() {
		return
	}
	syncVersionLockValue(c, r)
}

// zeroAllVersionLocks sets every versionLock in the tree to the value a create
// must carry.
func zeroAllVersionLocks(request any) {
	v := reflect.ValueOf(request)
	if !v.IsValid() {
		return
	}
	zeroVersionLockValue(v)
}

// topVersionLock reports the resource-level versionLock and whether one exists.
func topVersionLock(v any) (int, bool) {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return 0, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return 0, false
	}
	f := rv.FieldByName(versionLockFieldName)
	if !f.IsValid() || !isVersionLockInt(f.Kind()) {
		return 0, false
	}
	return int(f.Int()), true
}

func syncVersionLockValue(current, request reflect.Value) {
	for current.Kind() == reflect.Pointer || current.Kind() == reflect.Interface {
		if current.IsNil() || request.IsNil() {
			return
		}
		current, request = current.Elem(), request.Elem()
	}

	switch current.Kind() {
	case reflect.Struct:
		t := current.Type()
		for i := range t.NumField() {
			if t.Field(i).PkgPath != "" { // unexported
				continue
			}
			cf, rf := current.Field(i), request.Field(i)
			if t.Field(i).Name == versionLockFieldName && isVersionLockInt(cf.Kind()) {
				if rf.CanSet() {
					rf.SetInt(cf.Int())
				}
				continue
			}
			syncVersionLockValue(cf, rf)
		}
	case reflect.Slice, reflect.Array:
		// Position is the only correlation available; entries the request adds
		// beyond what the server returned keep their own values.
		n := min(current.Len(), request.Len())
		for i := range n {
			syncVersionLockValue(current.Index(i), request.Index(i))
		}
	}
}

func zeroVersionLockValue(v reflect.Value) {
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := range t.NumField() {
			if t.Field(i).PkgPath != "" {
				continue
			}
			f := v.Field(i)
			if t.Field(i).Name == versionLockFieldName && isVersionLockInt(f.Kind()) {
				if f.CanSet() {
					f.SetInt(newResourceVersionLock)
				}
				continue
			}
			zeroVersionLockValue(f)
		}
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			zeroVersionLockValue(v.Index(i))
		}
	}
}

func isVersionLockInt(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// apiErrorStatus reports the HTTP status carried by err, if it is an API error.
func apiErrorStatus(err error) (int, bool) {
	var apiErr *response.APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode, true
	}
	return 0, false
}

// isVersionLockConflict reports whether err is the server rejecting a write
// because the supplied lock was not current.
func isVersionLockConflict(err error) bool {
	status, ok := apiErrorStatus(err)
	return ok && (status == http.StatusConflict || status == http.StatusPreconditionFailed)
}

// updateWithVersionLock performs a version-locked update, handling the protocol
// so callers never supply a versionLock themselves.
//
// Each attempt reads current server state, copies every lock in the tree onto
// the request, and writes once.
//
// The write deliberately bypasses the HTTP client's retry logic. That layer
// treats PUT as idempotent, which is untrue here: the server consumes the lock
// on the first successful write, so a replayed body carries a spent token and
// is rejected — reporting a conflict for a write that succeeded. Retrying is
// only meaningful with a lock re-read from the server, which is what the next
// iteration of this loop does.
//
// These endpoints have also been observed returning 5xx on writes they fully
// applied. Such a failure is ambiguous, so after one the resource is re-read:
// a lock beyond the value submitted means this write was the one that moved it,
// and the refreshed resource is returned as a success. That check is skipped
// for conflicts, where the server has explicitly said it rejected the write —
// running it there would misread a competing writer's commit as this one's
// success and silently discard the caller's change.
func updateWithVersionLock[T any](
	c *Client,
	endpoint string,
	request *T,
	fetch func() (*T, error),
) (*T, error) {
	var lastErr error

	for range defaultVersionLockAttempts {
		current, err := fetch()
		if err != nil {
			return nil, err
		}

		submitted, hasLock := topVersionLock(current)
		syncAllVersionLocks(current, request)

		var result T
		resp, err := c.HTTP.DoRequestNoRetry("PUT", endpoint, request, &result)
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
		if err == nil {
			return &result, nil
		}
		lastErr = err

		if !isVersionLockConflict(err) {
			if hasLock {
				if after, ferr := fetch(); ferr == nil {
					if now, ok := topVersionLock(after); ok && now > submitted {
						return after, nil
					}
				}
			}
			// Nothing suggests a stale lock, so re-reading and resubmitting
			// would fail the same way.
			if !isRetryableVersionLockError(err) {
				return nil, err
			}
		}
		// Next iteration re-reads state, so the retry carries fresh locks.
	}

	return nil, fmt.Errorf("update to %s failed after %d attempts: %w", endpoint, defaultVersionLockAttempts, lastErr)
}

// isRetryableVersionLockError reports whether re-reading the locks and writing
// again could plausibly succeed. Server faults qualify; definite client errors
// (bad request, unauthorised, not found) fail identically whatever the lock.
func isRetryableVersionLockError(err error) bool {
	status, ok := apiErrorStatus(err)
	if !ok {
		return false
	}
	return status >= http.StatusInternalServerError
}
