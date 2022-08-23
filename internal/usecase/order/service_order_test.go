package usecase_test

import (
	"context"
	"errors"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase/order"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  interface{}
}

var mockOnlineOrderNumberRF = "123"
var mockOnlineOrderNumber = "100012912"
var errorFindOnlineOrderNumber = errors.New("order_mysql - FindOnlineOrderNumber - Mock Error")
var successResultGetRequestedOrder = createOnlineOrder()
var errorGetRequestedOrder = errors.New("order_mysql - GetRequestedOrder - Mock Error")

func serviceOrder(t *testing.T) (*usecase.OrderUseCase, *MockOrderMySQLRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockOrderMySQLRepo(mockCtl)

	legacyUseCase := usecase.NewOrderUseCase(repo)

	return legacyUseCase, repo
}

func Test_FindOnlineOrderNumber(t *testing.T) {
	t.Parallel()

	serviceOrder, repo := serviceOrder(t)

	tests := []test{
		{
			name: "success",
			mock: func() {
				repo.EXPECT().FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF).Return(&mockOnlineOrderNumber, nil)
			},
			res: &mockOnlineOrderNumber,
			err: nil,
		},
		{
			name: "error",
			mock: func() {
				repo.EXPECT().FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF).Return(nil, errorFindOnlineOrderNumber)
			},
			res: (*string)(nil),
			err: errorFindOnlineOrderNumber,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := serviceOrder.FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF)

			require.Equal(t, tc.res, res)
			require.Equal(t, tc.err, err)
		})
	}

}

func Test_GetRequestedOrder(t *testing.T) {
	t.Parallel()

	serviceOrder, repo := serviceOrder(t)

	tests := []test{
		{
			name: "success",
			mock: func() {
				repo.EXPECT().GetRequestedOrder(context.TODO(), mockOnlineOrderNumber).Return(successResultGetRequestedOrder, nil)
			},
			res: successResultGetRequestedOrder,
			err: nil,
		},
		{
			name: "error",
			mock: func() {
				repo.EXPECT().GetRequestedOrder(context.TODO(), mockOnlineOrderNumber).Return(nil, errorGetRequestedOrder)
			},
			res: (*entity.Order)(nil),
			err: errorGetRequestedOrder,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := serviceOrder.GetRequestedOrder(context.TODO(), mockOnlineOrderNumber)

			require.Equal(t, tc.res, res)
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
