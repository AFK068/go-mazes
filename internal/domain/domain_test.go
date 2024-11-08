package domain_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	parent := domain.NewCell(1, 1, nil)
	cell := domain.NewCell(2, 2, parent)

	assert.Equal(t, 2, cell.Row, "Expected row to be 2")
	assert.Equal(t, 2, cell.Col, "Expected col to be 2")
	assert.Equal(t, parent, cell.Parent, "Expected parent to be set correctly")
}

func TestGetChild(t *testing.T) {
	parent := domain.NewCell(1, 1, nil)
	cell := domain.NewCell(2, 2, parent)
	child := cell.GetChild()

	expectedChild := domain.NewCell(3, 3, cell)

	assert.Equal(t, expectedChild.Row, child.Row, "Expected child row to be 3")
	assert.Equal(t, expectedChild.Col, child.Col, "Expected child col to be 3")
	assert.Nil(t, child.Parent, "Expected child parent to be nil")
}
