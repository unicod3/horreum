package dbclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientImplementsDataStorer(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*DataStorage)(nil), new(Client))
}

func TestDataCollectionImplementsDataTable(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*DataTable)(nil), new(DataCollection))
}
