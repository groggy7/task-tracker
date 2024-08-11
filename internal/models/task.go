package models

type Task struct {
	ID          int
	Description string
	Done        bool
	Date        uint64
}
