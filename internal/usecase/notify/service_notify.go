package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/diogoalbuquerque/sub-notifier/pkg/net/client"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
)

const (
	_bodyType = "application/json"
)

type NotifyUsecase struct {
	client *client.RcClient
}

func NewNotifyUseCase(client *client.RcClient) *NotifyUsecase {
	return &NotifyUsecase{client: client}
}

func (uc *NotifyUsecase) NotifyPartner(ctx context.Context, url string, apiKey string, request interface{}) (*http.Response, error) {
	var (
		requestBody bytes.Buffer
	)

	if err := json.NewEncoder(&requestBody).Encode(request); err != nil {
		return nil, fmt.Errorf("service_notify - NotifyPartner - Encoder: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &requestBody)

	if err != nil {
		return nil, fmt.Errorf("service_notify - NotifyPartner - Request: %w", err)
	}

	req.Header.Set("Content-Type", _bodyType)
	req.Header.Set("Authorization", apiKey)

	resp, err := ctxhttp.Do(ctx, uc.client.Http, req)

	if err != nil {
		return nil, fmt.Errorf("service_notify - NotifyPartner - Post: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service_notify - NotifyPartner - Response: %d", resp.StatusCode)
	}

	return resp, nil
}
