package message

type Message struct {
	ID              int    `json:"id,omitempty"`
	ChatID          int    `json:"chat_id,omitempty"`
	ChatName        string `json:"chat_name,omitempty"`
	CreatorNickname string `json:"creator_nickname,omitempty"`
	TextMessage     string `json:"text_message,omitempty"`
}
