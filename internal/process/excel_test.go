package process

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/estenssoros/sheetdrop/models"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestExcel(t *testing.T) {
	home, _ := homedir.Dir()
	fileName := filepath.Join(home, "Downloads", "Crew JDE Mapping.xlsx")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	schema := &models.Schema{}
	assert.Nil(t, Excel(schema, data))
	fmt.Println(schema)
}
