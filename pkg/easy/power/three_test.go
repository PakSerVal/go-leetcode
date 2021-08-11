package power_test

import (
	"testing"

	"github.com/devbackend/go-leetcode/pkg/easy/power"
	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "example 1",
			num:      27,
			expected: true,
		},
		{
			name:     "example 2",
			num:      0,
			expected: false,
		},
		{
			name:     "example 3",
			num:      9,
			expected: true,
		},
		{
			name:     "example 4",
			num:      45,
			expected: false,
		},
		{
			name:     "example 5",
			num:      -9,
			expected: false,
		},
		{
			name:     "example 6",
			num:      1,
			expected: true,
		},
	}
	for _, c := range cases {
		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			actual := power.IsPowerOfThree(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
