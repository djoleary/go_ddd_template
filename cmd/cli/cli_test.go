package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

			err := run(mockIn, mockOut, mockErr, test.args)

			if test.expectErr == true {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
