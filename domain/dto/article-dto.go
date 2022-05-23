package dto

type ArticleUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

//BookCreateDTO is is a model that clinet use when create a new book
type ArticleCreateDTO struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Contetnt string `json:"content" form:"content" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
