package usecase

import (
	"testing"

	"github.com/yakan15/clean-transaction/store"
	"github.com/yakan15/clean-transaction/store/transactor"
)

type usecaseMocks struct {
	store      store.Store
	transactor transactor.Transactor
}

func TestCreate(t *testing.T) {
	// transactor だけ手動定義 mock を使用する
}
