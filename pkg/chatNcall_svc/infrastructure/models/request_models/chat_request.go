package requestmodels_chatNcallSvc_apigw

import "time"

type MessageRequest struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Content     string    `json:"Content" `
	TimeStamp   time.Time `json:"TimeStamp"`
	Type        string    `json:"Type" validate:"required"`
	Status      string    `json:"Status"`
	ChannelID   string    `json:"ChannelID"`
	ChatID      string    `json:"ChatID"`
	TypingStat  bool
}

type OnetoOneMessageRequest struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Type        string    `json:"Type"`
	Content     string    `json:"Content" validate:"required"`
	TimeStamp   time.Time `json:"TimeStamp"`
	Status      string    `json:"Status"`
}

type TypingStatusRequest struct {
	SenderID    string `json:"SenderID" `
	RecipientID string `json:"RecipientID"`
	Type        string `json:"Type" `
	TypingStat  bool
}
