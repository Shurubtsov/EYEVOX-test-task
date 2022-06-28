package message

type Message struct {
	ID              string `json:"id,omitempty"`
	CreatorNickname string `json:"creator_nickname,omitempty"`
	ChatID          string `json:"chat_id,omitempty"`
	TextMessage     string `json:"text_message,omitempty"`
}
