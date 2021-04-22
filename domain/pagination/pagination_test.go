package pagination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPaginatedList(t *testing.T) {
	pgn := NewPaginatedList(1 ,1, 10)
	assert.Equal(t, 1, pgn.Limit())

	pgn2 := NewPaginatedList(0, 0, 10)
	assert.Equal(t, 0, pgn2.Offset())
	assert.Equal(t, 15, pgn2.PerPage)
}

func TestPaginatedList_Limit(t *testing.T) {
	pgn := NewPaginatedList(1 ,5, 10)
	assert.Equal(t, 5, pgn.Limit())
}

func TestPaginatedList_Offset(t *testing.T) {
	pgn := NewPaginatedList(1 ,1, 10)
	assert.Equal(t, 0, pgn.Offset())
}

func TestGetPaginatedListFromRequest(t *testing.T) {
	// Not sure how to mock request
}
