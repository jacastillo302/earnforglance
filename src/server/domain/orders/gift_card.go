package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGiftCard = "gift_cards"
)

// GiftCard represents a gift card
type GiftCard struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	PurchasedWithOrderItemID *int               `bson:"purchased_with_order_item_id,omitempty"`
	GiftCardTypeID           int                `bson:"gift_card_type_id"`
	Amount                   float64            `bson:"amount"`
	IsGiftCardActivated      bool               `bson:"is_gift_card_activated"`
	GiftCardCouponCode       string             `bson:"gift_card_coupon_code"`
	RecipientName            string             `bson:"recipient_name"`
	RecipientEmail           string             `bson:"recipient_email"`
	SenderName               string             `bson:"sender_name"`
	SenderEmail              string             `bson:"sender_email"`
	Message                  string             `bson:"message"`
	IsRecipientNotified      bool               `bson:"is_recipient_notified"`
	CreatedOnUtc             time.Time          `bson:"created_on_utc"`
	GiftCardType             int                `bson:"gift_card_type"`
}
