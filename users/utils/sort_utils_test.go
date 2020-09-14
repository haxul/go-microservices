package utils

import (
	"github.com/stretchr/testify/assert"
	. "reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{10, 1, 9, 6}
	BubbleSort(list)
	equal := DeepEqual(list, []int{1, 6, 9, 10})
	assert.Equal(t, true, equal)
}
