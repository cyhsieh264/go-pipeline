package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client = New()
)

func TestFoo(t *testing.T) {
	got := client.Foo()
	assert.Equal(t, "bar", got)
}
