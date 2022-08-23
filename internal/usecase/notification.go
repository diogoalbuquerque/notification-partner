package usecase

import (
	"context"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/pkg/logger"
	"net/http"
)

type NotificationUseCase struct {
	sn ServiceNotify
	so ServiceOrder
	l  logger.Interface
}

func NewNotificationUseCase(sn ServiceNotify, so ServiceOrder, l logger.Interface) *NotificationUseCase {
	return &NotificationUseCase{
		sn: sn,
		so: so,
		l:  l,
	}
}

func (uc *NotificationUseCase) Notification(ctx context.Context, header string, url string, messageOrderRF string) error {
	var seg *xray.Segment

	var nrSeqOnlineOrder *string
	var err error

	err = xray.Capture(ctx, "FindOnlineOrderNumber", func(sc context.Context) error {
		nrSeqOnlineOrder, err = uc.so.FindOnlineOrderNumber(ctx, messageOrderRF)
		seg = xray.GetSegment(sc)
		seg.AddMetadata("Result", nrSeqOnlineOrder)
		seg.AddMetadata("Error", err)
		return err
	})

	if err != nil {
		seg.AddError(err)
		seg.Close(err)
		uc.l.Error(err)
		return err
	}

	var order *entity.Order

	err = xray.Capture(ctx, "GetRequestedOrder", func(sc context.Context) error {
		order, err = uc.so.GetRequestedOrder(ctx, *nrSeqOnlineOrder)
		seg = xray.GetSegment(sc)
		seg.AddMetadata("Result", order)
		seg.AddMetadata("Error", err)
		return err
	})

	if err != nil {
		seg.AddError(err)
		seg.Close(err)
		uc.l.Error(err)
		return err
	}

	var resp *http.Response

	err = xray.Capture(ctx, "NotifyPartner", func(sc context.Context) error {
		resp, err = uc.sn.NotifyPartner(ctx, url, header, entity.Notification{
			ClientRechargeToken:  order.Token,
			PaymentRequestNumber: messageOrderRF,
			ClientIdentity:       order.Cellphone,
		})
		seg = xray.GetSegment(sc)
		seg.AddMetadata("Result", resp)
		seg.AddMetadata("Error", err)
		return err
	})

	if err != nil {
		seg.AddError(err)
		seg.Close(err)
		uc.l.Error(err)
		return err
	}

	return nil

}
