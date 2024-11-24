package messages

type CreatePostRequest struct {
	PostID uint
	PostTitle string
	PostSubID uint
	PostBody string
	PostCreationTime string
	PostOwner uint
	PostKarma int
	Token string
}

type GetPostRequest struct {
	PostID uint
	PostTitle string
	PostSubID uint
	PostBody string
	PostOwner uint
	PostKarma int
	Token string
}

type CreatePostResponse struct {
	Error string
}

type DeletePostRequest struct {
	PostID uint
}
