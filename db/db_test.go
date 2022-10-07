package db_test

import (
	"main/db"
	"testing"
)

func TestPersistent(t *testing.T) {
	total := db.Openconnention()
	if total == nil {
		t.Errorf("Error Connection")
	}
}
