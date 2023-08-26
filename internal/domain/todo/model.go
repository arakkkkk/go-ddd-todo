package todo

import "time"

type Schema struct {
  Title string
  Completed bool
  Priority int
  CretedAt time.Time
}
