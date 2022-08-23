package repo_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/internal/usecase/order/repo"
	"github.com/diogoalbuquerque/sub-notifier/pkg/mysql"

	"github.com/stretchr/testify/assert"
	"testing"
)

var mockOnlineOrderNumberRF = "123"
var mockOnlineOrderNumber = "100012912"

func Test_FindOnlineOrderNumber_Success(t *testing.T) {

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.NR_SEQ_VENDA FROM VENDA_RF SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_RF = \\? FOR FETCH ONLY WITH UR;"

	rows := mock.NewRows(
		[]string{"NR_SEQ_VENDA_ONLINE"}).
		AddRow(mockOnlineOrderNumber)

	mock.ExpectQuery(query).WithArgs(mockOnlineOrderNumberRF).WillReturnRows(rows)

	onlineOrderNumber, err := r.FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF)

	assert.Equal(t, mockOnlineOrderNumber, *onlineOrderNumber, "The result should be the same.")
	assert.NoError(t, err, "This result should not have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func Test_FindOnlineOrderNumber_NoRowsError(t *testing.T) {

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.NR_SEQ_VENDA FROM VENDA_RF SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_RF = \\? FOR FETCH ONLY WITH UR;"

	mock.ExpectQuery(query).WithArgs(mockOnlineOrderNumberRF).WillReturnError(sql.ErrNoRows)

	onlineOrderNumber, err := r.FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF)

	assert.Nil(t, onlineOrderNumber, "The result should be nil.")
	assert.Error(t, err, "This result should have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func Test_FindOnlineOrderNumber_MockError(t *testing.T) {

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.NR_SEQ_VENDA FROM VENDA_RF SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_RF = \\? FOR FETCH ONLY WITH UR;"

	mock.ExpectQuery(query).WithArgs(mockOnlineOrderNumberRF).WillReturnError(errors.New("mock Error"))

	onlineOrderNumber, err := r.FindOnlineOrderNumber(context.TODO(), mockOnlineOrderNumberRF)

	assert.Nil(t, onlineOrderNumber, "The result should be nil.")
	assert.Error(t, err, "This result should have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func Test_GetRequestedOrder_Success(t *testing.T) {
	mockOrder := createOnlineOrder()

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.CD_SOLICITACAO, SQ_NOT_SUB.NR_CHIP_CARTAO_SC, SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE, SQ_NOT_SUB.TX_TOKEN, SQ_NOT_SUB.VL_CARGA, SQ_NOT_SUB.NR_CELL FROM SOLICITACAO_SUB SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE = \\? FOR FETCH ONLY WITH UR;"

	rows := mock.NewRows(
		[]string{"CD_SOLICITACAO_TEF48", "NR_CHIP_CARTAO_SC", "NR_SEQ_VENDA_ONLINE", "TX_TOKEN", "VL_CARGA", "NR_CELL_TOMAIS"}).
		AddRow(mockOrder.Code, mockOrder.NrChip, mockOrder.SeqOnlineOrder, mockOrder.Token, mockOrder.VlCharge, mockOrder.Cellphone)

	mock.ExpectQuery(query).WithArgs(mockOrder.SeqOnlineOrder).WillReturnRows(rows)

	order, err := r.GetRequestedOrder(context.TODO(), mockOrder.SeqOnlineOrder)

	assert.Equal(t, mockOrder.Code, order.Code, "The resul should be the same.")
	assert.Equal(t, mockOrder.NrChip, order.NrChip, "The resul should be the same.")
	assert.Equal(t, mockOrder.SeqOnlineOrder, order.SeqOnlineOrder, "The resul should be the same.")
	assert.Equal(t, mockOrder.Token, order.Token, "The resul should be the same.")
	assert.Equal(t, mockOrder.VlCharge, order.VlCharge, "The resul should be the same.")
	assert.Equal(t, mockOrder.Cellphone, order.Cellphone, "The resul should be the same.")
	assert.NoError(t, err, "This result should not have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func Test_GetRequestedOrder_NoRowsError(t *testing.T) {
	mockOrder := createOnlineOrder()

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.CD_SOLICITACAO, SQ_NOT_SUB.NR_CHIP_CARTAO_SC, SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE, SQ_NOT_SUB.TX_TOKEN, SQ_NOT_SUB.VL_CARGA, SQ_NOT_SUB.NR_CELL FROM SOLICITACAO_SUB SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE = \\? FOR FETCH ONLY WITH UR;"

	mock.ExpectQuery(query).WithArgs(mockOrder.SeqOnlineOrder).WillReturnError(sql.ErrNoRows)

	order, err := r.GetRequestedOrder(context.TODO(), mockOrder.SeqOnlineOrder)

	assert.Nil(t, order, "The result should be nil.")
	assert.Error(t, err, "This result should have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func Test_GetRequestedOrder_MockError(t *testing.T) {
	mockOrder := createOnlineOrder()

	connection, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("Error '%s' was not expected when opening a stub database connection", err))

	mysqlTest := createMySQLTest(connection)
	r := repo.NewOrderMySQLRepo(mysqlTest)

	defer r.Close()

	query := "SELECT SQ_NOT_SUB.CD_SOLICITACAO, SQ_NOT_SUB.NR_CHIP_CARTAO_SC, SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE, SQ_NOT_SUB.TX_TOKEN, SQ_NOT_SUB.VL_CARGA, SQ_NOT_SUB.NR_CELL FROM SOLICITACAO_SUB SQ_NOT_SUB WHERE SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE = \\? FOR FETCH ONLY WITH UR;"

	mock.ExpectQuery(query).WithArgs(mockOrder.SeqOnlineOrder).WillReturnError(errors.New("mock Error"))

	order, err := r.GetRequestedOrder(context.TODO(), mockOrder.SeqOnlineOrder)

	assert.Nil(t, order, "The result should be nil.")
	assert.Error(t, err, "This result should not have errors.")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("An error occurred and the following expectations were not met: %s", err)
	}

	assert.NoError(t, err, "This result should not have errors.")
}

func createMySQLTest(db *sql.DB) *mysql.MYSQL {
	return &mysql.MYSQL{
		MySQLDatabase: "Portal",
		DB:            db,
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
