package chat

type Chat struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FounderNickname string `json:"founder_nickname,omitempty"`
}

type CreateChatDTO struct {
	Name            string `json:"name,omitempty"`
	FounderNickname string `json:"founder_nickname,omitempty"`
}
