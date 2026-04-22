package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type FlexPayGateway struct {
	apiKey     string
	merchantID string
	baseURL    string
	httpClient *http.Client
}

func NewFlexPayGateway(apiKey, merchantID, baseURL string) ports.PaymentGateway {
	return &FlexPayGateway{
		apiKey:     apiKey,
		merchantID: merchantID,
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (g *FlexPayGateway) InitiatePayment(ctx context.Context, req ports.PaymentRequest) (*ports.PaymentResponse, error) {
	payload := map[string]interface{}{
		"merchant":    g.merchantID,
		"amount":      req.Amount.String(),
		"currency":    req.Currency,
		"phone":       req.PhoneNumber,
		"reference":   req.Reference,
		"callback":    req.CallbackURL,
		"description": req.Description,
	}
	body, _ := json.Marshal(payload)
	httpReq, _ := http.NewRequestWithContext(ctx, "POST", g.baseURL+"/payments", bytes.NewReader(body))
	httpReq.Header.Set("Authorization", "Bearer "+g.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse response...
	return &ports.PaymentResponse{TransactionID: "flexpay_tx_123", Status: "pending"}, nil
}
