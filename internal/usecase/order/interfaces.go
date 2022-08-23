package usecase

import (
	"context"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
)

type (
	OrderMySQLRepo interface {
		FindOnlineOrderNumber(ctx context.Context, seqOnlineOrderRf string) (*string, error)
		GetRequestedOrder(ctx context.Context, seqOnlineOrder string) (*entity.Order, error)
	}
)
