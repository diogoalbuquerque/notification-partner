package usecase_test

import (
	"context"
	"encoding/json"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase/notify"
	"github.com/diogoalbuquerque/sub-notifier/pkg/net/client"
	"github.com/stretchr/testify/assert"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	_mockDefaultApiKey  = "mock"
	_mockDefaultTimeout = 10
)

func Test_NotifyPartner_Success(t *testing.T) {
	mockResponse, _ := json.Marshal(createNotificationSuccessResponse())

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(mockResponse)

		}),
	)

	defer mockServer.Close()

	rcClient := client.New(client.MaxTimeout(_mockDefaultTimeout))
	rcClient.Http = mockServer.Client()

	resp, err := usecase.NewNotifyUseCase(rcClient).NotifyPartner(context.TODO(), mockServer.URL, _mockDefaultApiKey, createNotificationSuccessResponse())

	assert.NoError(t, err, "This result should not have errors.")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "The result should be the same.")

}

func Test_NotifyPartner_EncoderError(t *testing.T) {

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)

		}),
	)

	defer mockServer.Close()

	rcClient := client.New(client.MaxTimeout(_mockDefaultTimeout))
	rcClient.Http = mockServer.Client()

	_, err := usecase.NewNotifyUseCase(rcClient).NotifyPartner(context.TODO(), mockServer.URL, _mockDefaultApiKey, math.Inf(1))

	assert.Error(t, err, "This result should have errors.")

}

func Test_NotifyPartner_RequestError(t *testing.T) {

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)

		}),
	)

	defer mockServer.Close()

	rcClient := client.New(client.MaxTimeout(_mockDefaultTimeout))
	rcClient.Http = mockServer.Client()

	_, err := usecase.NewNotifyUseCase(rcClient).NotifyPartner(nil, mockServer.URL, _mockDefaultApiKey, nil)

	assert.Error(t, err, "This result should have errors.")

}

func Test_NotifyPartner_StatusCodeError(t *testing.T) {

	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)

		}),
	)

	defer mockServer.Close()

	rcClient := client.New(client.MaxTimeout(_mockDefaultTimeout))
	rcClient.Http = mockServer.Client()

	_, err := usecase.NewNotifyUseCase(rcClient).NotifyPartner(context.TODO(), mockServer.URL, _mockDefaultApiKey, createNotificationSuccessResponse())

	assert.Error(t, err, "This result should have errors.")

}

func createNotificationSuccessResponse() entity.Notification {
	return entity.Notification{
		ClientRechargeToken:  "f7278f1c-c001-4790-bec9-6908b1a7da40",
		PaymentRequestNumber: "123",
		ClientIdentity:       "21998876655",
	}
}
