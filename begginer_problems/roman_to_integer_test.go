package begginer_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRomanInteger struct {
	name           string
	input          string
	expectedResult int
}

func TestRomanToInt(t *testing.T) {

	tests := []testRomanInteger{
		{
			name:           "test 1",
			input:          "III",
			expectedResult: 3,
		},
		{
			name:           "test 2",
			input:          "VI",
			expectedResult: 6,
		},
		{
			name:           "test 3",
			input:          "IV",
			expectedResult: 4,
		},
		{
			name:           "test 4",
			input:          "MCMXCIV",
			expectedResult: 1994,
		},
		{
			name:           "test 5",
			input:          "MMMCMXCIX",
			expectedResult: 3999,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// my solution
			res := romanToInt(test.input)
			assert.Equal(t, test.expectedResult, res, test.name)
			// better solution
			resFastest := romanToIntFastest(test.input)
			assert.Equal(t, test.expectedResult, resFastest, test.name)
		})
	}
}
