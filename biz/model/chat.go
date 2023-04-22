package model

type MinimaxGPTReq struct {
	Model               string     `json:"model,required"`
	WithEmotion         bool       `json:"with_emotion,omitempty"`
	Stream              bool       `json:"stream,omitempty"`
	BeamWidth           bool       `json:"beam_width,omitempty"`
	Prompt              string     `json:"prompt"`
	Messages            []Messages `json:"messages,required"`
	ContinueLastMessage bool       `json:"continue_last_message,omitempty"`
	TokensToGenerate    int64      `json:"tokens_to_generate,omitempty"`
	Temperature         float32    `json:"temperature,omitempty"`
	TopP                float32    `json:"top_p,omitempty"`
	RoleMeta            `json:"role_meta"`
}

type RoleMeta struct {
	UserName string `json:"user_name"`
	BotName  string `json:"bot_name"`
}

type Messages struct {
	SenderType string `json:"sender_type"`
	Text       string `json:"text"`
}

type MinimaxGPTResp struct {
	Id              string    `json:"id"`
	Created         int64     `json:"created"`
	Model           string    `json:"model"`
	Reply           string    `json:"reply"`
	InputSensitive  bool      `json:"input_sensitive"`
	OutputSensitive bool      `json:"output_sensitive"`
	Choices         []Choices `json:"choices"`
	Usage           Usage     `json:"usage"`
	BaseResp        BaseResp  `json:"base_resp"`
}

type Choices struct {
	Text         string  `json:"text"`
	Index        int64   `json:"index"`
	LogProbes    float32 `json:"logprobes"`
	FinishReason string  `json:"finish_reason"`
	Emotion      string  `json:"emotion"`
	Delta        string  `json:"delta"`
}

type Usage struct {
	TotalTokens int64 `json:"total_tokens"`
}

type BaseResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
