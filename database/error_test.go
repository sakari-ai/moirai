package database

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestErrNotFound(t *testing.T) {
	var (
		cases = []struct {
			err      error
			expected bool
		}{
			{err: gorm.ErrRecordNotFound, expected: true},
			{err: ErrRecordNotFound, expected: true},
			{err: gorm.ErrInvalidSQL, expected: false},
			{err: nil, expected: false},
		}
	)
	for _, c := range cases {
		actual := ErrNotFound(c.err)
		assert.Equal(t, c.expected, actual, "case: {err: %s, expected: %v}", c.err, c.expected)
	}
}
