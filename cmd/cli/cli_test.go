package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEnv struct {
	pairs map[string]string
}

func (e testEnv) Getenv(k string) string {
	if v, ok := e.pairs[k]; ok {
		return v
	}
	return ""
}

func TestRun(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args     []string
		expected string
	}{
		"greet": {
			args:     []string{"greet"},
			expected: "hello world!\n",
		},
		"greet 'ted'": {
			args:     []string{"greet", "ted"},
			expected: "hello ted!\n",
		},
		"greet 'ted' 'bob'": {
			args:     []string{"greet", "ted", "bob"},
			expected: "hello ted and bob!\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockIn := new(bytes.Buffer)
			mockOut := new(bytes.Buffer)
			mockErr := new(bytes.Buffer)
			mockEnv := testEnv{}

			err := run(mockIn, mockOut, mockErr, mockEnv, test.args)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, mockOut.String())
		})
	}
}
