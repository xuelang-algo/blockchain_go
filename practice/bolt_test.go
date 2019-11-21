package practice

import (
	"log"
	"testing"

	"github.com/boltdb/bolt"
)

func TestDbOpening(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

