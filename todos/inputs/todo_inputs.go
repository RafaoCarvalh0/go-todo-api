package inputs

type CreateTodoInput struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
	Done  bool   `json:"done"`
}

type UpdateTodoInput struct {
	Title *string `json:"title,omitempty" validate:"min=3,max=255"`
	Done  *bool   `json:"done,omitempty"`
}
