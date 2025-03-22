package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGiftCard = "gift_cards"
)

// GiftCard represents a gift card
type GiftCard struct {
	ID                       primitive.ObjectID  `bson:"_id,omitempty"`
	PurchasedWithOrderItemID *primitive.ObjectID `bson:"purchased_with_order_item_id,omitempty"`
	GiftCardTypeID           int                 `bson:"gift_card_type_id"`
	Amount                   float64             `bson:"amount"`
	IsGiftCardActivated      bool                `bson:"is_gift_card_activated"`
	GiftCardCouponCode       string              `bson:"gift_card_coupon_code"`
	RecipientName            string              `bson:"recipient_name"`
	RecipientEmail           string              `bson:"recipient_email"`
	SenderName               string              `bson:"sender_name"`
	SenderEmail              string              `bson:"sender_email"`
	Message                  string              `bson:"message"`
	IsRecipientNotified      bool                `bson:"is_recipient_notified"`
	CreatedOnUtc             time.Time           `bson:"created_on_utc"`
	GiftCardType             int                 `bson:"gift_card_type"`
}

// GiftCardRepository interface
type GiftCardRepository interface {
	Create(c context.Context, gift_card *GiftCard) error
	Update(c context.Context, gift_card *GiftCard) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCard, error)
	FetchByID(c context.Context, ID string) (GiftCard, error)
}

// GiftCardUsecase interface
type GiftCardUsecase interface {
	FetchByID(c context.Context, ID string) (GiftCard, error)
	Create(c context.Context, gift_card *GiftCard) error
	Update(c context.Context, gift_card *GiftCard) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCard, error)
}
