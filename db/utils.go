package db

import (
  "log"
	"database/sql"
	"sync"
)

type Definition struct {
	db *sql.DB
	mu sync.Mutex
  id int
  name string
  description string
  phonetic string
  refs []int
  aliases []int
}

func AddDefinition(name string) bool {
  return false
}

func CheckError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
