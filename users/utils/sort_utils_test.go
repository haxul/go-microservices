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

func BenchmarkBubbleSort(b *testing.B) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort(list)
	}
}
