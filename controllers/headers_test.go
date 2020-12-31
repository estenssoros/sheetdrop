package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadersByIDs(t *testing.T) {
	out, err := ctl(t).HeadersByIDs([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestHeadersBySchemaIDs(t *testing.T) {
	out, err := ctl(t).HeadersBySchemaIDs([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
