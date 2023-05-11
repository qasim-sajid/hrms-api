package dbhandler

import "github.com/qasim-sajid/hrms-api/models"

// DbHandler specifies DB context
type DbHandler interface {
	SetupDB()
	CloseDB()

	AddEmployee(*models.Employee) (*models.Employee, int, error)
	GetAllEmployees() ([]*models.Employee, error)
	GetEmployeesWithFilters(searchParams map[string]interface{}) ([]*models.Employee, error)
	GetEmployee(userID int64) (*models.Employee, error)
	GetEmployeeWithIdentity(identity string) (*models.Employee, error)
	UpdateEmployee(userID int64, updates map[string]interface{}) (*models.Employee, error)
	DeleteEmployee(userID int64) error
	CheckEmployeeLogin(string, string) (*models.Employee, error)

	AddContract(*models.Contract) (*models.Contract, int, error)
	GetAllContracts() ([]*models.Contract, error)
	GetContractsWithFilters(searchParams map[string]interface{}) ([]*models.Contract, error)
	GetContract(contractID int64) (*models.Contract, error)
	UpdateContract(contractID int64, updates map[string]interface{}) (*models.Contract, error)
	DeleteContract(contractID int64) error

	AddTransaction(*models.Transaction) (*models.Transaction, int, error)
	GetAllTransactions() ([]*models.Transaction, error)
	GetTransactionsWithFilters(searchParams map[string]interface{}) ([]*models.Transaction, error)
	GetTransaction(transactionID int64) (*models.Transaction, error)
	UpdateTransaction(transactionID int64, updates map[string]interface{}) (*models.Transaction, error)
	DeleteTransaction(transactionID int64) error

	AddRequest(*models.Request) (*models.Request, int, error)
	GetAllRequests() ([]*models.Request, error)
	GetRequestsWithFilters(searchParams map[string]interface{}) ([]*models.Request, error)
	GetRequest(requestID int64) (*models.Request, error)
	UpdateRequest(requestID int64, updates map[string]interface{}) (*models.Request, error)
	DeleteRequest(requestID int64) error
}
