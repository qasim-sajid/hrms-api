package domain

import (
	"time"
)

// Request defines leave object
type Request struct {
	RequestId  int64     `json:"request_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	ActionAt   time.Time `json:"action_at"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	EmployeeID int64 `json:"employee_id"`
}
