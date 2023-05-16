package dbhandler

import "github.com/qasim-sajid/hrms-api/domain"

// DbHandler specifies DB context
type DbHandler interface {
	SetupDB()
	CloseDB()

	AddEmployee(*domain.Employee) (*domain.Employee, int, error)
	GetAllEmployees() ([]*domain.Employee, error)
	GetEmployeesWithFilters(searchParams map[string]interface{}) ([]*domain.Employee, error)
	GetEmployee(userID int64) (*domain.Employee, error)
	GetEmployeeWithIdentity(identity string) (*domain.Employee, error)
	UpdateEmployee(userID int64, updates map[string]interface{}) (*domain.Employee, error)
	DeleteEmployee(userID int64) error
	CheckEmployeeLogin(string, string) (*domain.Employee, error)

	AddContract(*domain.Contract) (*domain.Contract, int, error)
	GetAllContracts() ([]*domain.Contract, error)
	GetContractsWithFilters(searchParams map[string]interface{}) ([]*domain.Contract, error)
	GetContract(contractID int64) (*domain.Contract, error)
	UpdateContract(contractID int64, updates map[string]interface{}) (*domain.Contract, error)
	DeleteContract(contractID int64) error

	AddTransaction(*domain.Transaction) (*domain.Transaction, int, error)
	GetAllTransactions() ([]*domain.Transaction, error)
	GetTransactionsWithFilters(searchParams map[string]interface{}) ([]*domain.Transaction, error)
	GetTransaction(transactionID int64) (*domain.Transaction, error)
	UpdateTransaction(transactionID int64, updates map[string]interface{}) (*domain.Transaction, error)
	DeleteTransaction(transactionID int64) error

	AddRequest(*domain.Request) (*domain.Request, int, error)
	GetAllRequests() ([]*domain.Request, error)
	GetRequestsWithFilters(searchParams map[string]interface{}) ([]*domain.Request, error)
	GetRequest(requestID int64) (*domain.Request, error)
	UpdateRequest(requestID int64, updates map[string]interface{}) (*domain.Request, error)
	DeleteRequest(requestID int64) error
}
