package domain

import (
	"time"
)

// Contract defines leave object
type Contract struct {
	ContractId int64     `json:"contract_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	BasicPay   float32   `json:"basic_pay"`
	TotalPto   float32   `json:"total_pto"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	EmployeeID int64 `json:"employee_id"`
}
