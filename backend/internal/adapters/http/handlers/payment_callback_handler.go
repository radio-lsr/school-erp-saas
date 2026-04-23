package handlers

import (
    "net/http"
)

type PaymentCallbackHandler struct{}

func NewPaymentCallbackHandler() *PaymentCallbackHandler {
    return &PaymentCallbackHandler{}
}

func (h *PaymentCallbackHandler) FlexPayCallback(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}