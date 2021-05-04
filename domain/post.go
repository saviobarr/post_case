package domain

//Post represents an post entity
type Post struct {
	ID      int64   `json:"post_id"`
	Title   string  `json:"title"`
	Content float64 `json:"content"`
}
