package inputs

type CreatePostInput struct {
	Title   string `json:"title" binding:"required,min=3"`
	Content string `json:"content" binding:"required"`
}
