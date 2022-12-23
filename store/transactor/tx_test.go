package transactor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetTx(t *testing.T) {
	db := gorm.DB{}
	ctx := context.WithValue(context.TODO(), &txKey, &db)
	ret, ok := GetTx(ctx)
	assert.Equal(t, &db, ret)
	assert.True(t, ok)

	ret, ok = GetTx(context.TODO())
	assert.NotEqual(t, db, ret)
	assert.False(t, ok)
}
