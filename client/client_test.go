package client_test

import (
	"testing"

	"go-crdt-load-test/client"

	"github.com/stretchr/testify/assert"
)

func TestGetCount(t *testing.T) {
	count, err := client.GetCount()
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func TestInc(t *testing.T) {
	err := client.Inc()
	assert.Nil(t, err)
}
