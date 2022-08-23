package internal

import (
	"github.com/diogoalbuquerque/sub-notifier/config"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase"
	nuc "github.com/diogoalbuquerque/sub-notifier/internal/usecase/notify"
	ouc "github.com/diogoalbuquerque/sub-notifier/internal/usecase/order"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase/order/repo"
	"github.com/diogoalbuquerque/sub-notifier/pkg/logger"
	"github.com/diogoalbuquerque/sub-notifier/pkg/mysql"
	"github.com/diogoalbuquerque/sub-notifier/pkg/net/client"
	"github.com/diogoalbuquerque/sub-notifier/pkg/secret"

	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Run(ctx context.Context, cfg config.Config, awsSecret secret.AwsSecret, messages []events.SQSMessage) (*events.SQSEventResponse, error) {

	l := logger.New(cfg.Log.Level)

	mysqlDB, err := mysql.New(awsSecret, mysql.MaxIdleConns(cfg.MYSQL.IdleConnMax), mysql.MaxOpenConns(cfg.MYSQL.OpenConnMax), mysql.MaxLifetime(cfg.MYSQL.LifeConnMax))

	if err != nil {
		l.Error(err)
		return nil, err
	}

	defer mysqlDB.Close()

	or := repo.NewOrderMySQLRepo(mysqlDB)

	rcClient := client.New(client.MaxTimeout(cfg.RcClient.Timeout))

	var resp []events.SQSBatchItemFailure

	for _, message := range messages {

		result := usecase.NewNotificationUseCase(nuc.NewNotifyUseCase(rcClient), ouc.NewOrderUseCase(or), l).Notification(ctx, cfg.Partner.ApiKey, cfg.Partner.URL, message.Body)

		if result != nil {
			l.Error(fmt.Errorf("app - Run - ItemFailure - Arg: %s Error: %w", message.Body, result))
			resp = append(resp, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
		}

	}

	return &events.SQSEventResponse{BatchItemFailures: resp}, nil

}
