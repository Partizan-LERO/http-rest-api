package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("databaseURL")

	if databaseURL == "" {
		databaseURL = "host=localhost port=5436 user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev_test sslmode=disable"
	}

	os.Exit(m.Run())
}