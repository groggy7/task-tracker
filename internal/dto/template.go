package dto

type AddTmplReq struct {
	Name  string         `json:"name,omitempty"`
	Tasks []TemplateTask `json:"tasks,omitempty"`
}

type UpdateTmplReq struct {
	ID    int            `json:"id,omitempty"`
	Name  string         `json:"name,omitempty"`
	Tasks []TemplateTask `json:"tasks,omitempty"`
}

type DeleteTmplReq struct {
	ID int `json:"id,omitempty"`
}

type TemplateTask struct {
	TemplateID  int    `json:"template_id,omitempty"`
	Description string `json:"description,omitempty"`
}
