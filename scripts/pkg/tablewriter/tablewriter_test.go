package tablewriter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	table := NewTable()

	assert.NotNil(t, table)
}
