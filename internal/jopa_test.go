package internal

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		devCount int
		devRel   []int
		expected string
	}{
		{
			desc:     "1",
			devCount: 1,
			devRel:   []int{1000},
			expected: "Yes",
		},
		{
			desc:     "2",
			devCount: 2,
			devRel:   []int{1, 1},
			expected: "Yes",
		},
		{
			desc:     "3",
			devCount: 3,
			devRel:   []int{1, 1, 1},
			expected: "No",
		},
		{
			desc:     "4",
			devCount: 4,
			devRel:   []int{1, 1, 2, 2},
			expected: "Yes",
		},
		{
			desc:     "3-2",
			devCount: 3,
			devRel:   []int{1, 2, 1},
			expected: "Yes",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := jopa(tC.devCount, tC.devRel)
			if got != tC.expected {
				t.Errorf("Wanted %s, got %s", tC.expected, got)
			}
		})
	}
}
