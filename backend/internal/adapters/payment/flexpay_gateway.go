package payment

import (
    "context"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type FlexPayGateway struct {
    apiKey     string
    merchantID string
    baseURL    string
}

func NewFlexPayGateway(apiKey, merchantID, baseURL string) ports.PaymentGateway {
    return &FlexPayGateway{apiKey: apiKey, merchantID: merchantID, baseURL: baseURL}
}

func (g *FlexPayGateway) InitiatePayment(ctx context.Context, req ports.PaymentRequest) (*ports.PaymentResponse, error) {
    return &ports.PaymentResponse{TransactionID: "stub", Status: "pending"}, nil
}

func (g *FlexPayGateway) CheckPaymentStatus(ctx context.Context, transactionID string) (*ports.PaymentStatus, error) {
    return &ports.PaymentStatus{TransactionID: transactionID, Status: "success"}, nil
}