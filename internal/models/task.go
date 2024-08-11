package models

type Task struct {
	ID          int64
	Description string
	Done        bool
	Date        uint64
}
