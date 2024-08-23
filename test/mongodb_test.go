package test

import (
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestMongoDB(t *testing.T) {
	config.Load()
	db := database.NewMongoDB()
	assert.NotNil(t, db)
}
