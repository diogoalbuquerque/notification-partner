package repo

import (
	"context"
	"fmt"
	"github.com/diogoalbuquerque/sub-notifier/internal/entity"
	"github.com/diogoalbuquerque/sub-notifier/pkg/mysql"
)

type OrderMySQLRepo struct {
	*mysql.MYSQL
}

func NewOrderMySQLRepo(mysql *mysql.MYSQL) *OrderMySQLRepo {
	return &OrderMySQLRepo{mysql}
}

func (r *OrderMySQLRepo) FindOnlineOrderNumber(ctx context.Context, seqOnlineOrderRf string) (*string, error) {

	query := "SELECT " +
		"    SQ_NOT_SUB.NR_SEQ_VENDA " +
		"FROM VENDA_RF SQ_NOT_SUB " +
		"WHERE " +
		"    SQ_NOT_SUB.NR_SEQ_VENDA_RF = ? FOR FETCH ONLY WITH UR;"

	var seqOnlineOrder *string

	switch err := r.DB.QueryRowContext(ctx, query, seqOnlineOrderRf).Scan(&seqOnlineOrder); err {
	case nil:
		return seqOnlineOrder, err
	default:
		return nil, fmt.Errorf("order_mysql - FindOnlineOrderNumber - Query - Arg: %s Error: %w", seqOnlineOrderRf, err)
	}

}

func (r *OrderMySQLRepo) GetRequestedOrder(ctx context.Context, seqOnlineOrder string) (*entity.Order, error) {

	query := "SELECT " +
		"    SQ_NOT_SUB.CD_SOLICITACAO, " +
		"    SQ_NOT_SUB.NR_CHIP_CARTAO_SC, " +
		"    SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE, " +
		"    SQ_NOT_SUB.TX_TOKEN, " +
		"    SQ_NOT_SUB.VL_CARGA, " +
		"    SQ_NOT_SUB.NR_CELL " +
		"FROM SOLICITACAO_SUB SQ_NOT_SUB " +
		"WHERE " +
		"    SQ_NOT_SUB.NR_SEQ_VENDA_ONLINE = ? FOR FETCH ONLY WITH UR;"

	var order entity.Order

	switch err := r.DB.QueryRowContext(ctx, query, seqOnlineOrder).Scan(&order.Code, &order.NrChip, &order.SeqOnlineOrder, &order.Token, &order.VlCharge, &order.Cellphone); err {
	case nil:
		return &order, err
	default:
		return nil, fmt.Errorf("order_mysql - GetRequestedOrder - Query - Arg: %s Error: %w", seqOnlineOrder, err)
	}

}
