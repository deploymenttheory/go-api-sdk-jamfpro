// accounts_test.go
// TODO - 1st itteration
package jamfpro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountByID(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/JSSResource/accounts/userid/123", r.URL.String())
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id": 123, "name": "test", "password": "password"}`)
	}))
	defer mockServer.Close()

	cfg := Config{
		URL:        mockServer.URL,
		HTTPClient: mockServer.Client(),
		AuthMethod: BasicAuthConfig{Username: "your_username", Password: "your_password"}, // Use appropriate authentication method
	}

	client, err := NewClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	account, err := client.GetAccountByID(123)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, 123, account.ID)
	assert.Equal(t, "test", account.Name)
	assert.Equal(t, "password", account.Password)
}

func TestGetAccountByName(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/JSSResource/accounts/username/test", r.URL.String())
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id": 123, "name": "test", "password": "password"}`)
	}))
	defer mockServer.Close()

	cfg := Config{
		URL:        mockServer.URL,
		HTTPClient: mockServer.Client(),
		AuthMethod: BasicAuthConfig{Username: "your_username", Password: "your_password"}, // Use appropriate authentication method
	}

	client, err := NewClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	account, err := client.GetAccountByName("test")
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, 123, account.ID)
	assert.Equal(t, "test", account.Name)
	assert.Equal(t, "password", account.Password)
}

func TestGetAccounts(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/JSSResource/accounts", r.URL.String())
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"users": {"user": [{"id": 1, "name": "user1"}, {"id": 2, "name": "user2"}]}}`)
	}))
	defer mockServer.Close()

	cfg := Config{
		URL:        mockServer.URL,
		HTTPClient: mockServer.Client(),
		AuthMethod: BasicAuthConfig{Username: "your_username", Password: "your_password"}, // Use appropriate authentication method
	}

	client, err := NewClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	accountsList, err := client.GetAccounts()
	assert.NoError(t, err)
	assert.NotNil(t, accountsList)
	assert.Len(t, accountsList.Users.User, 2)
	assert.Equal(t, 1, accountsList.Users.User[0].ID)
	assert.Equal(t, "user1", accountsList.Users.User[0].Name)
	assert.Equal(t, 2, accountsList.Users.User[1].ID)
	assert.Equal(t, "user2", accountsList.Users.User[1].Name)
}

// Similarly, you can write tests for other functions like CreateAccount, UpdateAccountByID, DeleteAccountByName, etc.
