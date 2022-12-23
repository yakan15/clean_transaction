package transactor

import (
	"context"

	"gorm.io/gorm"
)

type Transactor interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error)
}

type mockTransactor struct{}

func NewMockTransactor() Transactor {
	return &mockTransactor{}
}

func (m *mockTransactor) DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	return f(ctx)
}

var txKey = struct{}{}

type tx struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) *tx {
	return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	tx := t.db.Begin()
	ctx = context.WithValue(ctx, &txKey, tx)

	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
