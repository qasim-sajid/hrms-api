package domain

import (
	"time"
)

// Transaction defines leave object
type Transaction struct {
	TransactionId   int64     `json:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date"`
	PaidAmount      float32   `json:"paid_amount"`
	AvailedPto      float32   `json:"availed_pto"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	ContractId int64 `json:"contract_id"`
}
