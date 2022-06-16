package hotspot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	// Arrange
	cli := cliStub{config: func() ([]byte, error) {
		return []byte("/vendor/\n/something/\n"), nil
	}}

	// Act
	config, err := parseConfig(cli)

	// Assert
	assert.Nil(t, err, "failed to parse config")
	assert.Len(t, config, 2)
}

type cliStub struct {
	config func() ([]byte, error)
}

func (c cliStub) Config() ([]byte, error) {
	return c.config()
}

func (c cliStub) Files() ([]byte, error) {
	l := "/vendor/\n/something/"
	return []byte(l), nil
}
