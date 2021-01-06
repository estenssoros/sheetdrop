package controllers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/go-homedir"
	"github.com/satori/uuid"
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

func TestCreateDeleteSchemaExcel(t *testing.T) {
	home, _ := homedir.Dir()
	fileName := "all-terminals.xlsx"
	f, err := os.Open(filepath.Join(home, "Downloads", fileName))
	if err != nil {
		t.Skip("could not find file")
	}
	out, err := ctl(t).CreateSchema(&CreateSchemaInput{
		ResourceID: 1,
		Name:       uuid.NewV1().String(),
		File: &graphql.Upload{
			File:     f,
			Filename: fileName,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "excel", out.SourceType)
	assert.Equal(t, fileName, out.SourceURI)
	_, err = ctl(t).DeleteSchemaByID(out.ID)
}

func TestCreateDeleteSchemaCSV(t *testing.T) {
	home, _ := homedir.Dir()
	fileName := "all-terminals.csv"
	f, err := os.Open(filepath.Join(home, "Downloads", fileName))
	if err != nil {
		t.Skip("could not find file")
	}
	out, err := ctl(t).CreateSchema(&CreateSchemaInput{
		ResourceID: 1,
		Name:       uuid.NewV1().String(),
		File: &graphql.Upload{
			File:     f,
			Filename: fileName,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "csv", out.SourceType)
	assert.Equal(t, fileName, out.SourceURI)
	_, err = ctl(t).DeleteSchemaByID(out.ID)
}
