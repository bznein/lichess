package board

type ChatMessage struct {
	Text string `json:"text"`
	User string `json:"user"`
}

type Chat []ChatMessage
