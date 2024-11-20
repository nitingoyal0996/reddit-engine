package messages

type KarmaRequest struct {
	UserId 	uint
	Amount 	int
	Token 	string
}

type KarmaResponse struct {
	Error 	string
}