package dto

type AddTmplReq struct {
	Name  string           `json:"name,omitempty"`
	Tasks []TmplTaskAddReq `json:"tasks,omitempty"`
}

type UpdateTmplReq struct {
	ID    int                 `json:"id,omitempty"`
	Name  string              `json:"name,omitempty"`
	Tasks []TmplTaskUpdateReq `json:"tasks,omitempty"`
}

type DeleteTmplReq struct {
	ID int `json:"id,omitempty"`
}

type TmplTaskAddReq struct {
	Description string `json:"description,omitempty"`
}

type TmplTaskUpdateReq struct {
	ID          int    `json:"id,omitempty"`
	TemplateID  int    `json:"template_id,omitempty"`
	Description string `json:"description,omitempty"`
}
