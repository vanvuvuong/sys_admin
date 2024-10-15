package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleMethod(t *testing.T) {
	t.Parallel()
	circle := &Circle{1.2}
	circle.bigger()
	assert.Equal(t, float32(12), circle.Radius, "Circle should be multiple.")
}
