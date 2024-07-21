package models

import (
	"fmt"
	"time"
)

const TransactionTypeBuy = "BUY"
const TransactionTypeTopUp = "TOP_UP"
const TransactionTypeGift = "GIFT"

type Transaction struct {
	Base
	Type            string    `gorm:"type:varchar(10);not null" json:"type"`
	FromID          string    `gorm:"type:uuid;not null" json:"from_id"`
	ToID            string    `gorm:"type:varchar(100)" json:"to_id"`
	ProductID       string    `gorm:"type:varchar(100)" json:"product_id"`
	Price           int64     `gorm:"default:0" json:"price"`
	TransactionDate time.Time `json:"transaction_date"`
}

func (t *Transaction) GetPriceFormatted() string {
	b := float64(t.Price) / 100.00
	return fmt.Sprintf("%.2f", b)
}
