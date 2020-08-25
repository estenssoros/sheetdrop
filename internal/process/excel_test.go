package process

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestExcel(t *testing.T) {
	home, _ := homedir.Dir()
	fileName := filepath.Join(home, "Downloads", "Crew JDE Mapping.xlsx")
	stat, err := os.Stat(fileName)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	out, err := Excel(f, stat.Size())
	assert.Nil(t, err)
	fmt.Println(out)
}
