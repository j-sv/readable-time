package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Clock(t *testing.T) {
	tests := []struct {
		timestamp time.Time
		expected  string
	}{
		{
			timestamp: time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC),
			expected:  "noon",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 2, 0, 0, time.UTC),
			expected:  "noon",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 11, 58, 0, 0, time.UTC),
			expected:  "noon",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 13, 0, 0, 0, time.UTC),
			expected:  "one o'clock",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 3, 0, 0, time.UTC),
			expected:  "five past twelve",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 11, 53, 0, 0, time.UTC),
			expected:  "five to twelve",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected:  "midnight",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 15, 0, 0, time.UTC),
			expected:  "quarter past twelve",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 45, 0, 0, time.UTC),
			expected:  "quarter to one",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 25, 0, 0, time.UTC),
			expected:  "twenty-five past twelve",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 35, 0, 0, time.UTC),
			expected:  "twenty-five to one",
		},
		{
			timestamp: time.Date(2020, time.January, 1, 12, 30, 0, 0, time.UTC),
			expected:  "half past twelve",
		},
	}

	for _, test := range tests {
		t.Run(test.timestamp.Format(time.RFC3339), func(t *testing.T) {
			actual := Time{test.timestamp}.Clock()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_Day(t *testing.T) {
	tests := []struct {
		timestamp time.Time
		expected  string
	}{
		{
			timestamp: time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC),
			expected:  "first",
		},
		{
			timestamp: time.Date(2020, time.January, 11, 12, 0, 0, 0, time.UTC),
			expected:  "eleventh",
		},
		{
			timestamp: time.Date(2020, time.January, 18, 12, 0, 0, 0, time.UTC),
			expected:  "eighteenth",
		},
		{
			timestamp: time.Date(2020, time.January, 21, 12, 0, 0, 0, time.UTC),
			expected:  "twenty-first",
		},
		{
			timestamp: time.Date(2020, time.January, 30, 12, 0, 0, 0, time.UTC),
			expected:  "thirtieth",
		},
	}

	for _, test := range tests {
		t.Run(test.timestamp.Format(time.RFC3339), func(t *testing.T) {
			actual := Time{test.timestamp}.Day()
			assert.Equal(t, test.expected, actual)
		})
	}
}
