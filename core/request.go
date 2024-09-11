package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)

type Client struct {
	apiURL     string
	referURL   string
	authToken  string
	httpClient *http.Client
}

func handleResponse(respBody []byte) (map[string]interface{}, error) {
	// Mengurai JSON ke dalam map[string]interface{}
	var result map[string]interface{}
	err := json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return result, nil
}

func (c *Client) makeRequest(method string, endpoint string, jsonBody interface{}) ([]byte, error) {
	fullURL := c.apiURL + endpoint

	// Convert body to JSON
	var reqBody []byte
	var err error
	if jsonBody != nil {
		reqBody, err = json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
	}

	// Create new request
	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set header
	setHeader(req, c.referURL, c.authToken)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 status code
	if resp.StatusCode >= 400 {
		// Read the response body to include in the error message
		bodyBytes, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			return nil, fmt.Errorf("error status: %v, and failed to read body: %v", resp.StatusCode, bodyErr)
		}
		return nil, fmt.Errorf("error status: %v, error message: %s", resp.StatusCode, string(bodyBytes))
	}

	return io.ReadAll(resp.Body)
}

// Login Account
func (c *Client) login(queryData string) ([]byte, error) {
	payload := map[string]string{
		"tgData": queryData,
	}

	return c.makeRequest("POST", "/api/uxp/user/login", payload)
}

// Get Detail Account
func (c *Client) detailAccount() ([]byte, error) {
	payload := map[string]string{}

	return c.makeRequest("POST", "/api/uxp/user/detail", payload)
}

// Daily Check In
func (c *Client) dailyCheckIn(queryData string) ([]byte, error) {
	payload := map[string]string{
		"tgData": queryData,
	}

	return c.makeRequest("POST", "/api/uxp/user/checkIn", payload)
}

// Claim Task
func (c *Client) claimTask(taskType string) ([]byte, error) {
	payload := map[string]string{}

	return c.makeRequest("POST", fmt.Sprintf("/api/uxp/user/%s", taskType), payload)
}

// Bind Wallet
func (c *Client) bindWallet(walletAddress string) ([]byte, error) {
	payload := map[string]string{
		"wallet": walletAddress,
	}

	return c.makeRequest("POST", "/api/uxp/user/bindWallet", payload)
}
