package usecase

import (
	"context"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
)

type OrderUseCase struct {
	repo OrderMySQLRepo
}

func NewOrderUseCase(repo OrderMySQLRepo) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}

func (uc *OrderUseCase) FindOnlineOrderNumber(ctx context.Context, seqOnlineOrderRf string) (*string, error) {
	return uc.repo.FindOnlineOrderNumber(ctx, seqOnlineOrderRf)
}

func (uc *OrderUseCase) GetRequestedOrder(ctx context.Context, seqOnlineOrder string) (*entity.Order, error) {
	return uc.repo.GetRequestedOrder(ctx, seqOnlineOrder)
}
