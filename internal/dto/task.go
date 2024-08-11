package dto

type AddTaskReq struct {
	Description string `json:"description,omitempty"`
}

type AddFromTmplReq struct {
	TemplateID int `json:"template_id,omitempty"`
}

type UpdateTaskReq struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done,omitempty"`
}

type DeleteTaskReq struct {
	ID int `json:"id,omitempty"`
}
