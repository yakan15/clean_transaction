package store

import (
	"context"

	"github.com/yakan15/clean-transaction/model"
	"github.com/yakan15/clean-transaction/store/transactor"
	"gorm.io/gorm"
)

type Store interface {
	Create(ctx context.Context, a *model.Article) error
	Update(ctx context.Context, a *model.Article) error
}

type store struct {
	d *gorm.DB
}

func NewStore(d *gorm.DB) *store {
	return &store{d: d}
}

func (s *store) Create(ctx context.Context, a *model.Article) error {
	d, ok := transactor.GetTx(ctx)
	if !ok {
		d = s.d
	}
	if err := d.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (s *store) Update(ctx context.Context, a *model.Article) error {
	d, ok := transactor.GetTx(ctx)
	if !ok {
		d = s.d
	}
	if err := d.Save(a).Error; err != nil {
		return err
	}
	return nil
}
