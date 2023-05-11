package models

import "time"

// Employee defines employee object
type Employee struct {
	EmployeeID   int64     `json:"employee_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	EmployeeType string    `json:"employee_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	ManagerID int64 `json:"manager_id"`
}
