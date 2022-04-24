package dash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackEmpty(t *testing.T) {
	assert := assert.New(t)

	s := newStackReverse([]string{"value"})
	assert.False(s.empty())
	s.pop()
	assert.True(s.empty())
}

func TestStackPop(t *testing.T) {
	assert := assert.New(t)

	s := newStackReverse([]string{"foo", "bar", "baz"})

	assert.Equal("baz", s.pop())
	assert.Equal("bar", s.pop())
	assert.Equal("foo", s.pop())
	assert.Panics(
		func() { s.pop() },
		"should panic on 4th pop on stack with 3 items",
	)
	s.push("a")
	s.push("b")
	assert.Equal("b", s.pop())
	assert.Equal("a", s.pop())
	assert.Panics(
		func() { s.pop() },
		"should panic due to empty",
	)
}

func TestStackPush(t *testing.T) {
	assert := assert.New(t)

	s := newStackReverse([]string{"foo", "bar", "baz"})

	assert.Panics(
		func() { s.push("a") },
		"should panic due to over capacity",
	)

	assert.Equal("baz", s.pop())
	assert.Equal("bar", s.pop())

	s.push("a")
	s.push("b")

	assert.Panics(
		func() { s.push("c") },
		"should panic due to over capacity",
	)
}
