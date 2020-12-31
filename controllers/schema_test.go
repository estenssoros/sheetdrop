package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaHeaders(t *testing.T) {
	out, err := ctl(t).SchemaHeaders(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}

func TestSchemaByID(t *testing.T) {
	out, err := ctl(t).SchemaByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestSchemasForResource(t *testing.T) {
	out, err := ctl(t).SchemasForResource(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestListSchemas(t *testing.T) {
	out, err := ctl(t).ListSchemas()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
func TestSchemasByIDs(t *testing.T) {
	out, err := ctl(t).SchemasByIDs([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, out)
}
