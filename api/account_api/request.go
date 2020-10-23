package account_api

type GetAccountRequest struct {
	Ex         string `bson:"ex" json:"ex" binding:"required"`
	TickerType string `bson:"ticker_type" json:"ticker_type" binding:"required"`
	Ticker     string `b