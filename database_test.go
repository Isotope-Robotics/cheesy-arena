// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package main

import (
	"os"
	"testing"
)

const testDbPath = "test.db"

func clearDb() {
	os.Remove(testDbPath)
}

func TestOpenUnreachableDatabase(t *testing.T) {
	_, err := OpenDatabase("nonexistentdir/test.db")
	if err == nil {
		t.Errorf("Expected error not raised")
	}
}
