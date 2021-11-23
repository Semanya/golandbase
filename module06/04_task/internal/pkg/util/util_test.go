package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testpair struct {
	check   []int
	currect bool
}

var tests = []testpair{
	{[]int{123, 321, 331231, 2321}, false},
	{[]int{-649, -946, -233121, -3213}, false},
	{[]int{-64, 46, 3121, 213}, false},
}

func TestContainsDuplicate(t *testing.T) {
	req := require.New(t)
	for _, pair := range tests {
		v := ContainsDuplicate(pair.check)
		req.Equal(v, pair.currect)
	}

}

func TestIsPalindrome(t *testing.T) {
	req := require.New(t)
	PalCase := func(number int, res bool) func(t *testing.T) {
		return func(t *testing.T) {
			res2 := IsPalindrome(number)
			req.Equal(res2, res)
		}
	}
	t.Run("simple case", PalCase(355, false))
	t.Run("zero values", PalCase(533, false))
	t.Run("negative values", PalCase(725, false))
}
