package secret

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/diogoalbuquerque/sub-notifier/config"
)

type secret struct {
	ctx    context.Context
	config config.Config
}

type Secret interface {
	Load() (*AwsSecret, error)
}

type AwsSecret struct {
	MySQLUsername string `json:"mysql_username,omitempty"`
	MySQLPassword string `json:"mysql_password,omitempty"`
	MySQLEngine   string `json:"mysql_engine,omitempty"`
	MySQLHost     string `json:"mysql_host,omitempty"`
	MySQLPort     string `json:"mysql_port,omitempty"`
	MySQLDatabase string `json:"mysql_database,omitempty"`
}

func New(ctx context.Context, config config.Config) Secret {
	return &secret{ctx: ctx,
		config: config}
}

func (s *secret) Load() (*AwsSecret, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(s.config.RegionName)})
	svc := secretsmanager.New(sess)

	xray.AWS(svc.Client)

	if err != nil {
		return nil, fmt.Errorf("secret - Load - NewSession: %w", err)
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(s.config.SecretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValueWithContext(s.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("secret - Load - GetSecretValueWithContext: %w", err)
	}

	secret := &AwsSecret{}

	err = json.Unmarshal([]byte(*result.SecretString), &secret)
	if err != nil {
		return nil, fmt.Errorf("secret - Load - Unmarshal: %w", err)
	}

	return secret, nil
}
