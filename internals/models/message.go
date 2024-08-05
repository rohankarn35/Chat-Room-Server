package models

type Message struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	Content     string `json:"content"`
	IsJoined    bool   `json:false`
}
