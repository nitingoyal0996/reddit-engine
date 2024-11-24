package messages

type CreateCommentRequest struct {
	CommentID uint
	CommentBody string
	CommentPostID uint
	CommentOwner uint
	CommentKarma int
	CommentParentID uint
	CommentCreationTime string
	CommentIsDeleted bool
	Token string
}

type GetCommentRequest struct {
	CommentID uint
	CommentBody string
	CommentPostID uint
	CommentOwner uint
	CommentKarma int
	CommentParentID uint
	CommentIsDeleted bool
}

type DeleteCommentRequest struct {
	CommentID uint
	Token string
}	

type CreateCommentResponse struct {
	Error string
}