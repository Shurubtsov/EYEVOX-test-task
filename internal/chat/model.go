package chat

type Chat struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FounderNickname string `json:"founder_nickname,omitempty"`
}
