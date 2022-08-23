# ⚙️ notification-partner

[![Build](https://github.com/diogoalbuquerque/notification-partner/actions/workflows/build.yml/badge.svg)](https://github.com/diogoalbuquerque/notification-partner/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/diogoalbuquerque/notification-partner/branch/main/graph/badge.svg?token=KsP5u5ljhu)](https://codecov.io/gh/diogoalbuquerque/notification-partner)

## Project

This project consists of a payment notification queue. When a payment for an order is successfully made,
an object is placed in the queue for processing and sending the notification to the partner.

The flow starts when an order is paid.

After payment of the order, another system will send the number of the paid order to the queue.

When the order is added to the queue, this function will start.

- Each added order will have its existence verified in the customer order table;
- Then a notification will be sent to the partner's api informing that the order has been released and is ready to be
  used.

### Usage

Requires an sqs object to start.

## Technologies used

This function was built using the following technologies:

- [aws-lambda-go - v1.34.1](https://github.com/aws/aws-lambda-go);
- [aws-sdk-go - v1.44.82](https://github.com/aws/aws-sdk-go);
- [aws-xray-sdk-go - v1.7.0](https://github.com/aws/aws-xray-sdk-go);
- [cleanenv v1.3.0](https://github.com/ilyakaznacheev/cleanenv);
- [go-sqlmock - v1.5.0](https://github.com/DATA-DOG/go-sqlmock);
- [mysql - v1.6.0](https://github.com/go-sql-driver/mysql);
- [testify v1.8.0](https://github.com/stretchr/testify);
- [zerolog v1.27.0](https://github.com/rs/zerolog).

## Technical operation

### How to install

## Main variables to be defined

| Variable            | Variable description                               | Default value                                   |
|---------------------|----------------------------------------------------|-------------------------------------------------|
| LOG_LEVEL           | Log level.                                         | info                                            |
| MYSQL_OPEN_CONN_MAX | Maximum number of connections that can be opened.  | 100                                             |
| MYSQL_IDLE_CONN_MAX | Maximum number of connections that can be stopped. | 2                                               |
| MYSQL_LIFE_CONN_MAX | Maximum connection time in seconds.                | 10                                              |
| AWS_REGION_NAME     | Region where the secrets manager is located.       | DEFINE_REGION_NAME                              |
| SECRETS_MANAGER     | Name of the secret manager used in the application | SECRETS_MANAGER                                 |
| RC_CLIENT_TIMEOUT   | Default timeout for http client.              | 30                                              |
| PARTNER_API_KEY     | Authorization key for partner api.                 | Key cmlvY2FyZHJvdXRlcjpHdXZHTUNHM3              |
| PARTNER_URL         | Partner notification address.                      | https://hmg-cs.mock.ai/api/Recharge/pix-payment |

#### Project build:

To upload a function you need to create and build it like linux operating.
So use the following command in the project folder:

```sh
$ GOOS=linux GOARCH=amd64 go build -o SUB_NOTIFIER ./
```

or you can use Makefile.

```sh
$ make build
```

❗ Attention️

- After the build, you need to create the .zip of the generated file and then upload this .zip file;
- The Handler is the name of the generated .zip file.

## Comments

- This project is a part of a proof of concept used in the company, so parts of them were omitted as well as domain and
  data model information.