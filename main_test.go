package main_test

import (
	"ci-cd/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {

	assert.Equal(t, "Hello World!", handler.SayHello())

}
