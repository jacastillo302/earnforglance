package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionGiftCard = "gift_cards"
)

// GiftCard represents a gift card
type GiftCard struct {
	ID                  bson.ObjectID  `bson:"_id,omitempty"`
	OrderItemID         *bson.ObjectID `bson:"purchased_with_order_item_id"`
	GiftCardTypeID      int            `bson:"gift_card_type_id"`
	Amount              float64        `bson:"amount"`
	IsGiftCardActivated bool           `bson:"is_gift_card_activated"`
	GiftCardCouponCode  string         `bson:"gift_card_coupon_code"`
	RecipientName       string         `bson:"recipient_name"`
	RecipientEmail      string         `bson:"recipient_email"`
	SenderName          string         `bson:"sender_name"`
	SenderEmail         string         `bson:"sender_email"`
	Message             string         `bson:"message"`
	IsRecipientNotified bool           `bson:"is_recipient_notified"`
	CreatedOnUtc        time.Time      `bson:"created_on_utc"`
}

// GiftCardRepository interface
type GiftCardRepository interface {
	CreateMany(c context.Context, items []GiftCard) error
	Create(c context.Context, gift_card *GiftCard) error
	Update(c context.Context, gift_card *GiftCard) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCard, error)
	FetchByID(c context.Context, ID string) (GiftCard, error)
}

// GiftCardUsecase interface
type GiftCardUsecase interface {
	CreateMany(c context.Context, items []GiftCard) error
	FetchByID(c context.Context, ID string) (GiftCard, error)
	Create(c context.Context, gift_card *GiftCard) error
	Update(c context.Context, gift_card *GiftCard) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCard, error)
}
