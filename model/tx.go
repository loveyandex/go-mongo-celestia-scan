package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TiaTxn struct {
	ID2 primitive.ObjectID `bson:"_id,omitempty" json:"bid,omitempty"`

	ID            int       `json:"id"`
	Height        int       `json:"height"`
	Position      int       `json:"position"`
	GasWanted     int       `json:"gas_wanted"`
	GasUsed       int       `json:"gas_used"`
	TimeoutHeight int       `json:"timeout_height"`
	EventsCount   int       `json:"events_count"`
	MessagesCount int       `json:"messages_count"`
	Hash          string    `json:"hash"`
	Fee           string    `json:"fee"`
	Time          time.Time `json:"time"`
	MessageTypes  []string  `json:"message_types"`
	Status        string    `json:"status"`
	Memo          string    `json:"memo,omitempty"`
}
type EventTxn struct {
	ID2 primitive.ObjectID `bson:"_id,omitempty" json:"bid"`
	
	ID       int         `json:"id"`
	Height   int         `json:"height"`
	Time     time.Time   `json:"time"`
	Position int         `json:"position"`
	TxID     int         `json:"tx_id"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}

type TransferEvent struct {
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
}
type EventType string

const (
	TransferEventType = "transfer"
)