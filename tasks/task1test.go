package tasks_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"asia_yo_take_home_task/tasks"
)

func TestFindTheMinNum(t *testing.T) {
	tests := []struct {
		description   string
		input         []int
		wantMin       int
		wantFrequency int
	}{
		{
			description:   "test case 1",
			input:         []int{1, 9, 15, 3, 29, 19},
			wantMin:       1,
			wantFrequency: 1,
		},
		{
			description:   "test case 2",
			input:         []int{4999, 3999, 2999, 1999},
			wantMin:       1999,
			wantFrequency: 1,
		},
		{
			description:   "test case 3",
			input:         []int{4999, 3999, 2999, 1999, 1999, 1999},
			wantMin:       1999,
			wantFrequency: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			gotMin, gotFrequency := tasks.FindTheMinNum(tt.input)
			require.Equal(t, tt.wantMin, gotMin)
			require.Equal(t, tt.wantFrequency, gotFrequency)
		})
	}
}
