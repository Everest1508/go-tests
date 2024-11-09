package integration

import (
	"net/http"
	"testing"

	"github.com/everest1508/go-tests-setup/config"
	apiClient "github.com/everest1508/go-tests-setup/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestLoginAPI(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	if err != nil {
		t.Fatal("Cannot load configurations:", err)
	}

	client := apiClient.NewAPIClient(config.API.BaseURL)

	endpoint := "/core/login"
	payload := map[string]string{
		"email":    "ritesh@scalent.io",
		"password": "Scalent@123",
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := client.Post(endpoint, payload, headers)
	if err != nil {
		t.Fatalf("Failed to perform POST request: %v", err)
	}

	assert.Equal(t, response.StatusCode, http.StatusOK, "failed because status code")
}

func TestSum(t *testing.T) {
	a := 2
	b := 10
	c := a + b
	assert.Equal(t, c, 10, `not equal bro`)
}
