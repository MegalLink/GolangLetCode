package begginer_problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testlongestCommonPrefix struct {
	name           string
	input          []string
	expectedResult string
}

func TestLongestCommonPrefixInt(t *testing.T) {
	tests := []testlongestCommonPrefix{
		{
			name:           "test 1",
			input:          []string{"flower", "flow", "flight"},
			expectedResult: "fl",
		},
		{
			name:           "test 2",
			input:          []string{"flower", "flow", "flight", "car"},
			expectedResult: "",
		},
		{
			name:           "test 3",
			input:          []string{"dog", "racecar", "car"},
			expectedResult: "",
		},
		{
			name:           "test 4",
			input:          []string{"ab", "a"},
			expectedResult: "a",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// my solution
			res := longestCommonPrefix(test.input)
			assert.Equal(t, test.expectedResult, res, test.name)
			resFaster := longestCommonPrefixFaster(test.input)
			assert.Equal(t, test.expectedResult, resFaster, test.name)
		})
	}
}
