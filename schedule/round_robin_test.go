package schedule_test

import (
	"testing"

	"go-crdt-load-test/schedule"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	const host1 = "a"
	const host2 = "b"
	const host3 = "c"
	rr := schedule.NewRoundRobin([]string{host1, host2, host3})

	assert.Equal(t, host1, rr.Next())
	assert.Equal(t, host2, rr.Next())
	assert.Equal(t, host3, rr.Next())
	assert.Equal(t, host1, rr.Next())
	assert.Equal(t, host2, rr.Next())
	assert.Equal(t, host3, rr.Next())
}
