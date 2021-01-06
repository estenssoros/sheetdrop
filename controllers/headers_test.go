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

func TestHeaderByID(t *testing.T) {
	out, err := ctl(t).HeaderByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestHeaderSetID(t *testing.T) {
	_, err := ctl(t).SetHeaderID(1, true)
	if err != nil {
		t.Fatal(err)
	}
	header, err := ctl(t).HeaderByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, header.IsID)
}
