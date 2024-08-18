package test

import (
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	config.Load()
	db := database.NewPostgreSQL()
	assert.NotNil(t, db)
}
