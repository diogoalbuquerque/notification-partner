package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diogoalbuquerque/sub-notifier/config"
	"github.com/diogoalbuquerque/sub-notifier/internal"
	"github.com/diogoalbuquerque/sub-notifier/pkg/secret"
	"github.com/rs/zerolog/log"
)

func syncNotificationPixHandle(ctx context.Context, sqsEvent events.SQSEvent) (*events.SQSEventResponse, error) {
	c, err := config.NewConfig()

	if err != nil {
		log.Err(err)
		return nil, err
	}

	s, err := secret.New(ctx, *c).Load()

	if err != nil {
		log.Err(err)
		return nil, err
	}

	return internal.Run(ctx, *c, *s, sqsEvent.Records)
}

func main() {
	lambda.Start(syncNotificationPixHandle)
}
