package quantum

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Quantum struct {
	Ctx        context.Context
	SqlConn    sqlx.SqlConn
	QuantumKey string

	GetQuantumKey interface{}
}

type GetQuantumKey interface {
	GetQuantumKey(ctx context.Context) (string, error)
}

func (qb *Quantum) GetQuantumKeyLogic(ctx context.Context) (string, error) {
	return "实现qubit逻辑", nil
}
