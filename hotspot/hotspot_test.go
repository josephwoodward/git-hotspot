package hotspot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	config, err := parseConfig(cliStub{})

	assert.Nil(t, err, "failed to parse config")
	assert.Len(t, config, 2)
}

type cliStub struct{}

func (c cliStub) Config() ([]byte, error) {
	l := "/vendor/\n/something/"
	return []byte(l), nil
}

func (c cliStub) Files() ([]byte, error) {
	l := "/vendor/\n/something/"
	return []byte(l), nil
}
