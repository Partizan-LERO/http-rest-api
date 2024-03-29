package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()
	config  := NewConfig()
	config.DatabaseURL = databaseUrl
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE ", strings.Join(tables, ", ")));
			err != nil {
				t.Fatal(err)
			}
			defer s.Close()
		}
	}
}