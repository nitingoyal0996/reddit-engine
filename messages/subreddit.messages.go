package messages

type GetSubredditRequest struct {
	SubID uint `json:"sub_id"`
	SubName string `json:"sub_name"`
	SubDesc string `json:"sub_desc"`
	SubCreationTime string `json:"sub_creation_time"`
	SubOwner string `json:"sub_owner"`
	Token string `json:"token"`
}

type CreateSubredditRequest struct {
	SubID uint `json:"sub_id"`
	SubName string `json:"sub_name"`
	SubDesc string `json:"sub_desc"`
	SubCreationTime string `json:"sub_creation_time"`
	SubOwner string `json:"sub_owner"`
	Token string `json:"token"`
}

type CreateSubredditResponse struct {
	Error string `json:"error"`
}