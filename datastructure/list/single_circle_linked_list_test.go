package list

import (
	"testing"
)

func TestNewSingleCircleLinkedList(t *testing.T) {

	list := NewSingleCircleLinkedList()
	testList(t, list)
}
