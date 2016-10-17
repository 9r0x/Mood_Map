package cognitive

type TextAnalytics struct {
	BingKey string
}

type TextDetail struct {
	ID    string  `json:"id"`
	Score float32 `json:"score"`
}

type TextReq struct {
	Language string `json:"language"`
	ID       string `json:"id"`
	Text     string `json:"text"`
}

type TextRequest struct {
	Documents []TextReq `json:"documents"`
}

type TextResult struct {
	Documents []TextDetail `json:"documents"`
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}
