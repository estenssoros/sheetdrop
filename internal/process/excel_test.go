package process

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

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
	out, err := Excel(data)
	assert.Nil(t, err)
	fmt.Println(out)
}
