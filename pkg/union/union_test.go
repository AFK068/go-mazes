package union_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/union"
	"github.com/stretchr/testify/assert"
)

func TestNewUnionFindSet(t *testing.T) {
	set := union.NewUnionFindSet(1)

	assert.Equal(t, 1, set.Parent, "Expected parent to be 1")
	assert.Equal(t, 0, set.Rank, "Expected rank to be 0")
}

func TestFind(t *testing.T) {
	sets := make([]*union.FindSet, 3)
	for i := range sets {
		sets[i] = union.NewUnionFindSet(i)
	}

	assert.Equal(t, 0, union.Find(sets, 0), "Expected Find(0) to return 0")
	assert.Equal(t, 1, union.Find(sets, 1), "Expected Find(1) to return 1")
	assert.Equal(t, 2, union.Find(sets, 2), "Expected Find(2) to return 2")

	union.Join(sets, 0, 1)
	assert.Equal(t, union.Find(sets, 0), union.Find(sets, 1), "Expected Find(0) and Find(1) to return the same root")
}

func TestJoin(t *testing.T) {
	sets := make([]*union.FindSet, 3)
	for i := range sets {
		sets[i] = union.NewUnionFindSet(i)
	}

	union.Join(sets, 0, 1)
	assert.Equal(t, union.Find(sets, 0), union.Find(sets, 1), "Expected Find(0) and Find(1) to return the same root")

	union.Join(sets, 1, 2)
	assert.Equal(t, union.Find(sets, 0), union.Find(sets, 2), "Expected Find(0) and Find(2) to return the same root")
	assert.Equal(t, union.Find(sets, 1), union.Find(sets, 2), "Expected Find(1) and Find(2) to return the same root")
}
