package datetime_test

import (
	"errors"
	"go-exercise/datetime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDateEndOfYear(t *testing.T) {

	type input struct {
		year int
	}

	type expected struct {
		result string
		err    error
	}

	type testCase struct {
		name     string
		input    input
		expected expected
	}

	testCases := []testCase{
		{
			name: "success date_end_of_year_2024",
			input: input{
				year: 2024,
			},
			expected: expected{
				result: "2024-12-31",
				err:    nil,
			},
		},
		{
			name: "success date_end_of_year_2025",
			input: input{
				year: 2025,
			},
			expected: expected{
				result: "2025-12-31",
				err:    nil,
			},
		},
		{
			name: "error date_end_of_year_20 ",
			input: input{
				year: 20,
			},
			expected: expected{
				result: "",
				err:    errors.New(`parsing time \"20-01-02\" as \"2006-01-02\": cannot parse \"20-01-02\" as \"2006\"`),
			},
		},
	}

	for _, testCase := range testCases {
		actual, err := datetime.GetDateEndOfYear(testCase.input.year)
		if testCase.expected.err != nil {
			assert.EqualError(t, err, testCase.expected.err.Error())
		}

		assert.Equal(t, testCase.expected, actual)
	}
}
