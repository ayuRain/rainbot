package model

type VerifyReq struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

type VerifyResp struct {
	Challenge string `json:"challenge"`
}
