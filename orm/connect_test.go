package orm

import "testing"

func TestConnect(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
}
