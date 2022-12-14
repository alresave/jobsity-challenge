package service

type StockRequest struct {
	StockCode string `json:"stockCode"`
	Room      string `json:"room"`
	UserName  string `json:"userName"`
}

type ChatRequest struct {
	UserName string `json:"userName"`
	Room     string `json:"room"`
	Message  string `json:"message"`
	DateTime int64  `json:"dateTime"`
}

type ConnectRequest struct {
	UserName string `json:"userName"`
	Room     string `json:"room"`
}
