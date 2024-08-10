package models

type Template struct {
	ID    int
	Name  string
	Tasks []TemplateTask
}

type TemplateTask struct {
	ID          int
	TemplateID  int
	Description string
}
