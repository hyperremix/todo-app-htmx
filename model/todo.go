package model

type Todo struct {
	ID          int64  `param:"id" query:"id" form:"id" json:"id"`
	Title       string `param:"title" query:"title" form:"title" json:"title"`
	Description string `param:"description" query:"description" form:"description" json:"description"`
	IsCompleted bool   `param:"is_completed" query:"is_completed" form:"is_completed" json:"is_completed"`
}
