package minStack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Problem(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	ast.Equal(-3, s.GetMin(), "get min from [-2, 0, -3]")
	s.Pop()
	ast.Equal(0, s.Top(), "get top from [-2, 0]")
	ast.Equal(-2, s.GetMin(), "get min from [-2, 0]")

}
