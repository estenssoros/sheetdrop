package process

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/estenssoros/sheetdrop/models"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestExcel(t *testing.T) {
	home, _ := homedir.Dir()
	fileName := filepath.Join(home, "Downloads", "Crew JDE Mapping.xlsx")
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	schema := &models.Schema{}
	_, err = Excel(schema, f)
	assert.Nil(t, err)
	fmt.Println(schema)
}
