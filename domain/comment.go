package domain

//Comment represents an comment entity
type Comment struct {
	ID      int64   `json:"comment_id"`
	Post    Post    `json:"post"`
	Content float64 `json:"content"`
}
