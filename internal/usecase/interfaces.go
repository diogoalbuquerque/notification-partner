package usecase

import (
	"context"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"net/http"
)

type (
	ServiceNotify interface {
		NotifyPartner(ctx context.Context, url string, header string, request interface{}) (*http.Response, error)
	}

	ServiceOrder interface {
		FindOnlineOrderNumber(ctx context.Context, seqOnlineOrderRf string) (*string, error)
		GetRequestedOrder(ctx context.Context, seqOnlineOrder string) (*entity.Order, error)
	}
)
