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
		cmd       string
		args      []string
		expectErr bool
	}{
		"empty echo": {
			cmd:       "echo",
			args:      []string{"echo"},
			expectErr: true,
		},
		"echo 'hello world'": {
			cmd:       "echo",
			args:      []string{"echo", "hello", "world"},
			expectErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockIn := new(bytes.Buffer)
			mockOut := new(bytes.Buffer)
			mockErr := new(bytes.Buffer)
			mockEnv := testEnv{}

			err := run(mockIn, mockOut, mockErr, mockEnv, test.args)

			if test.expectErr == true {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
