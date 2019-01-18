package triangle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	triangle [][]int
}

type ans struct {
	one int
}

func Test_Problem0120(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{

				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			}},
			ans{
				11,
			},
		},

		question{
			para{[][]int{}},
			ans{
				0,
			},
		},

		question{
			para{[][]int{
				[]int{2},
			}},
			ans{
				2,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, minimumTotal(p.triangle), "输入:%v", p)
	}
}