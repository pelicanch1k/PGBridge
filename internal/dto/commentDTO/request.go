package commentDTO

type CreateCommentRequestDTO struct {
	UserID int
	Params map[string]interface{}
}

type GetCommentRequestDTO struct {
	CommentID int
}

type UpdateCommentRequestDTO struct {
	CommentID int
	Params map[string]interface{}
}

type DeleteCommentRequestDTO struct {
	CommentID int
}
