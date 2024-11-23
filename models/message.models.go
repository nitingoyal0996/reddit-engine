package models

type Message struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	FromId    uint    `json:"fromId"`
	ToId      uint    `json:"toId"`
	CreatedAt string `json:"createdAt"`
}
