package usecase_test

import (
	"context"
	"errors"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase"
	"github.com/diogoalbuquerque/sub-notifier/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	_mockDefaultHeader = "mock"
	_mockDefaultUrl    = "https://mockurl.com"
	_mockLogLevel      = "error"
)

var mockOnlineOrderNumberRF = "123"
var mockOnlineOrderNumber = "100012912"
var errorFindOnlineOrderNumber = errors.New("order_mysql - FindOnlineOrderNumber - Mock Error")
var successResultGetRequestedOrder = createOnlineOrder()
var errorGetRequestedOrder = errors.New("order_mysql - GetRequestedOrder - Mock Error")
var successResultNotifyPartner = createNotificationSuccess()
var errorNotifyPartner = errors.New("service_notify - NotifyPartner - Mock Error")

type test struct {
	name string
	mock func()
	err  interface{}
}

func serviceNotification(t *testing.T) (*usecase.NotificationUseCase, *MockServiceNotify, *MockServiceOrder) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	serviceNotify := NewMockServiceNotify(mockCtl)
	serviceOrder := NewMockServiceOrder(mockCtl)

	notificationUseCase := usecase.NewNotificationUseCase(serviceNotify, serviceOrder, logger.New(_mockLogLevel))

	return notificationUseCase, serviceNotify, serviceOrder
}

func Test_Notification(t *testing.T) {
	t.Parallel()
	ctx, _ := xray.BeginSegment(context.TODO(), "CONTEXT_TEST")
	notification, serviceNotify, serviceOrder := serviceNotification(t)

	tests := []test{
		{
			name: "no rows find online order number error",
			mock: func() {
				serviceOrder.EXPECT().FindOnlineOrderNumber(ctx, mockOnlineOrderNumberRF).Return(nil, errorFindOnlineOrderNumber)
			},
			err: errorFindOnlineOrderNumber,
		},
		{
			name: "no rows get requested order error",
			mock: func() {
				gomock.InOrder(
					serviceOrder.EXPECT().FindOnlineOrderNumber(ctx, mockOnlineOrderNumberRF).Return(&mockOnlineOrderNumber, nil),
					serviceOrder.EXPECT().GetRequestedOrder(ctx, mockOnlineOrderNumber).Return(nil, errorGetRequestedOrder))
			},
			err: errorGetRequestedOrder,
		},
		{
			name: "notify partner error",
			mock: func() {
				gomock.InOrder(
					serviceOrder.EXPECT().FindOnlineOrderNumber(ctx, mockOnlineOrderNumberRF).Return(&mockOnlineOrderNumber, nil),
					serviceOrder.EXPECT().GetRequestedOrder(ctx, mockOnlineOrderNumber).Return(successResultGetRequestedOrder, nil),
					serviceNotify.EXPECT().NotifyPartner(ctx, _mockDefaultUrl, _mockDefaultHeader, successResultNotifyPartner).Return(nil, errorNotifyPartner))
			},
			err: errorNotifyPartner,
		},
		{
			name: "success",
			mock: func() {
				gomock.InOrder(
					serviceOrder.EXPECT().FindOnlineOrderNumber(ctx, mockOnlineOrderNumberRF).Return(&mockOnlineOrderNumber, nil),
					serviceOrder.EXPECT().GetRequestedOrder(ctx, mockOnlineOrderNumber).Return(successResultGetRequestedOrder, nil),
					serviceNotify.EXPECT().NotifyPartner(ctx, _mockDefaultUrl, _mockDefaultHeader, successResultNotifyPartner).Return(nil, nil))
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			err := notification.Notification(ctx, _mockDefaultHeader, _mockDefaultUrl, mockOnlineOrderNumberRF)

			require.Equal(t, tc.err, err)

		})
	}
}

func createOnlineOrder() *entity.Order {
	return &entity.Order{
		Code:           32,
		NrChip:         3670549064,
		SeqOnlineOrder: mockOnlineOrderNumber,
		Token:          "f7278f1c-c001-4790-bec9-6908b1a7da40",
		VlCharge:       60.00,
		Cellphone:      "21998876655",
	}
}

func createNotificationSuccess() entity.Notification {
	return entity.Notification{
		ClientRechargeToken:  "f7278f1c-c001-4790-bec9-6908b1a7da40",
		PaymentRequestNumber: mockOnlineOrderNumberRF,
		ClientIdentity:       "21998876655",
	}
}
