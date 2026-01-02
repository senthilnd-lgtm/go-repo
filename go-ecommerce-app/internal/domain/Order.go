package domain

import "time"

type Order struct {
	ID             uint        `json:"id" gorm:"PrimaryKey"`
	UserId         uint        `json:"user_id"`
	Status         string      `json:"status" `
	Amount         float64     `json:"amount"`
	TransactionId  string      `json:"transaction_id"`
	OrderRefNumber uint        `json:"order_ref_number"`
	PaymentId      string      `json:"payment_id"`
	Item           []OrderItem `json:"items"`
	CreatedAt      time.Time   `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt      time.Time   `json:"updated_at" gorm:"default:current_timestamp"`
}
