package game

import (
	"fmt"
	"testing"
)

func TestGameDateTic(t *testing.T) {
	var d GameDate = 0
	d.tic()
	if d != 10 {
		t.Errorf("date = %d, expected 10", d)
	}
}

func TestGetFormatDate(t *testing.T) {
	testCases := []struct {
		desc   string
		input  int
		output struct {
			m int
			h int
			d int
		}
	}{
		{
			desc:  "Zero hours",
			input: 30,
			output: struct {
				m int
				h int
				d int
			}{
				m: 30,
				h: 0,
				d: 0,
			},
		},
		{
			desc:  "Zero days",
			input: 605,
			output: struct {
				m int
				h int
				d int
			}{
				m: 4,
				h: 4,
				d: 1,
			},
		},
		{
			desc:  "Full string",
			input: 6010,
			output: struct {
				m int
				h int
				d int
			}{
				m: 5,
				h: 10,
				d: 0,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var dt GameDate = GameDate(tC.input)
			got := dt.GetFormatDate()
			want := fmt.Sprintf("Day: %d, %02d:%02d", tC.output.d, tC.output.h, tC.output.m)

			if want != got {
				t.Errorf("DATE = {%s}, WANT {%s}", got, want)
			}
		})
	}
}
