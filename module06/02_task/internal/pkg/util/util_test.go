package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testpair struct {
	values int
	result int
}

var tests = []testpair{
	{123, 321},
	{-649, -946},
	{0, 0},
	{333, 333},
}

func TestReverseInt(t *testing.T) {
	for _, pair := range tests {
		v, err := ReverseInt(pair.values)
		if err != nil {
			t.Errorf("fail to get orders: %w", err)
		}
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func Test2ReverseInt(t *testing.T) {
	req := require.New(t)
	for _, pair := range tests {
		v, err := ReverseInt(pair.values)
		req.NoError(err)
		req.Equal(v, pair.result)
	}

}
