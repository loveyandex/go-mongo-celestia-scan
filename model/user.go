package model

import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	USER       Role = "USER"
	ADMIN      Role = "ADMIN"
	SUPERADMIN Role = "SUPERADMIN"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	PassWord string `bson:"password,omitempty" json:"password,omitempty"`
	OTPCode  string `bson:"otp_code,omitempty" json:"otp_code,omitempty"`
	Roles    []Role `bson:"roles,omitempty" json:"roles,omitempty"`

	UserStatus  UserState  `bson:"user_status" json:"user_status"`
	TelUser     *TelUser   `bson:"tel_user" json:"tel_user,omitempty"`
 
	Transfering *Transfering 
	Wallets     []Wallet            `json:"wallets,omitempty" bson:"wallets,omitempty"` 
}

type TelUser struct {
	// ID is a unique identifier for this user or bot
	ID int64 `json:"id"`
	// IsBot true, if this user is a bot
	//
	// optional
	IsBot bool `json:"is_bot,omitempty"`
	// FirstName user's or bot's first name
	FirstName string `json:"first_name"`
	// LastName user's or bot's last name
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// UserName user's or bot's username
	//
	// optional
	UserName string `json:"username,omitempty"`
	// LanguageCode IETF language tag of the user's language
	// more info: https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// optional
	LanguageCode string `json:"language_code,omitempty"`
	// CanJoinGroups is true, if the bot can be invited to groups.
	// Returned only in getMe.
	//
	// optional
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// CanReadAllGroupMessages is true, if privacy mode is disabled for the bot.
	// Returned only in getMe.
	//
	// optional
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// SupportsInlineQueries is true, if the bot supports inline queries.
	// Returned only in getMe.
	//
	// optional
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

type UserState int

const (
	ZERO UserState = iota
	Buying
	PuttingBuyingAmount
	VerifiedOrder
	GOTOPAY

	TransferingInChatInline
	TransferingTo
	TransferingWalletType
	TransferingWalletValue
	TransferingConfirm
	TransferingFinished
	TransferingWalletValueIsOver
)

type Transfering struct {
	ToUser   primitive.ObjectID
	FromUser primitive.ObjectID
	Symbol   string
	Amount   float64
}
type Wallet struct {
	Name               string  `json:"name"`
	TokenId            string  `json:"token_id" bson:"token_id"`
	FullName           string  `json:"full_name"`
	PublicKey          string  `json:"public_key"`
	Value              float64 `json:"value"`
	ActivatedOnChain   bool    `json:"activated_on_chain" bson:"activated_on_chain"`
	UserOnChainBalance float64 `json:"on_chain_balance" bson:"on_chain_balance"`
	OffChainBalance    float64 `json:"off_chain_balance" bson:"off_chain_balance"`
	BrokerChainBalance float64 `json:"broker_chain_balance" bson:"broker_chain_balance"`
	Lock               float64
}
