package messages

type UpdateKarmaRequest struct {
	UserId 	uint
	Amount 	int
	Token 	string
}

type KarmaRequest struct {
	UserId 	uint
	Amount 	int
}

type KarmaResponse struct {
	Error 	string
}