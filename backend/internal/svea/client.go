package svea

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/config"
)

// Client is a client for the Svea Ekonomi API
type Client struct {
	httpClient *http.Client
	config     config.SveaConfig
}

// NewClient creates a new Svea client
func NewClient(config config.SveaConfig) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		config: config,
	}
}

// CreateOrder creates a new order in Svea Ekonomi
func (c *Client) CreateOrder(ctx context.Context, order *OrderRequest) (*OrderResponse, error) {
	// Prepare request URL
	url := fmt.Sprintf("%s/api/orders", c.config.BaseURL)

	// Prepare request body
	body, err := json.Marshal(order)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal order request: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Merchant-Id", c.config.MerchantID)
	req.Header.Set("X-Secret-Token", c.config.Secret)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var orderResponse OrderResponse
	if err := json.Unmarshal(respBody, &orderResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &orderResponse, nil
}

// GetOrder gets an order from Svea Ekonomi
func (c *Client) GetOrder(ctx context.Context, orderID string) (*Order, error) {
	// Prepare request URL
	url := fmt.Sprintf("%s/api/orders/%s", c.config.BaseURL, orderID)

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Merchant-Id", c.config.MerchantID)
	req.Header.Set("X-Secret-Token", c.config.Secret)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var order Order
	if err := json.Unmarshal(respBody, &order); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &order, nil
}

// FinalizePayment finalizes a payment in Svea Ekonomi
func (c *Client) FinalizePayment(ctx context.Context, orderID string, paymentMethod string) (*PaymentResponse, error) {
	// Prepare request URL
	url := fmt.Sprintf("%s/api/orders/%s/payments", c.config.BaseURL, orderID)

	// Prepare request body
	paymentRequest := PaymentRequest{
		PaymentMethod: paymentMethod,
	}
	body, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Merchant-Id", c.config.MerchantID)
	req.Header.Set("X-Secret-Token", c.config.Secret)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var paymentResponse PaymentResponse
	if err := json.Unmarshal(respBody, &paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &paymentResponse, nil
}

// CancelOrder cancels an order in Svea Ekonomi
func (c *Client) CancelOrder(ctx context.Context, orderID string) error {
	// Prepare request URL
	url := fmt.Sprintf("%s/api/orders/%s/cancel", c.config.BaseURL, orderID)

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Merchant-Id", c.config.MerchantID)
	req.Header.Set("X-Secret-Token", c.config.Secret)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// RefundPayment refunds a payment in Svea Ekonomi
func (c *Client) RefundPayment(ctx context.Context, orderID string, refundRequest *RefundRequest) (*RefundResponse, error) {
	// Prepare request URL
	url := fmt.Sprintf("%s/api/orders/%s/refunds", c.config.BaseURL, orderID)

	// Prepare request body
	body, err := json.Marshal(refundRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal refund request: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Merchant-Id", c.config.MerchantID)
	req.Header.Set("X-Secret-Token", c.config.Secret)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var refundResponse RefundResponse
	if err := json.Unmarshal(respBody, &refundResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &refundResponse, nil
}