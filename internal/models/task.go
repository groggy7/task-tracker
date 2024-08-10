package models

import (
	"time"
)

type Task struct {
	ID          int64
	Description string
	Done        bool
	Date        time.Time
}
