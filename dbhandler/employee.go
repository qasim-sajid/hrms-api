package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/qasim-sajid/hrms-api/domain"
)

func (db *dbClient) AddEmployee(Employee *domain.Employee) (*domain.Employee, int, error) {
	if status, err := db.CheckForDuplicateEmployee(Employee.Email); err != nil {
		return nil, status, fmt.Errorf("AddEmployee: %v", err)
	}

	if status, err := db.CheckForDuplicateEmployee(Employee.Username); err != nil {
		return nil, status, fmt.Errorf("AddEmployee: %v", err)
	}

	insertQuery, err := db.GetInsertQuery(*Employee)
	if err != nil {
		return nil, -1, fmt.Errorf("AddEmployee: %v", err)
	}

	_, err = db.RunInsertQuery(insertQuery)
	if err != nil {
		return nil, -1, fmt.Errorf("AddEmployee: %v", err)
	}

	return Employee, http.StatusOK, nil
}

func (db *dbClient) GetAllEmployees() ([]*domain.Employee, error) {
	Employees, err := db.GetEmployeesWithFilters(make(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("GetAllEmployees: %v", err)
	}

	return Employees, nil
}

func (db *dbClient) GetEmployee(EmployeeID int64) (*domain.Employee, error) {
	selectParams := make(map[string]interface{})

	selectParams["employee_id"] = EmployeeID

	Employees, err := db.GetEmployeesWithFilters(selectParams)
	if err != nil {
		return nil, fmt.Errorf("GetEmployee: %v", err)
	}

	var Employee *domain.Employee
	if Employees == nil || len(Employees) <= 0 {
		return nil, nil //fmt.Errorf("GetEmployee: %v", errors.New("Employee with given ID not found!"))
	} else {
		Employee = Employees[0]
	}

	return Employee, nil
}

func (db *dbClient) GetEmployeeWithIdentity(identity string) (*domain.Employee, error) {
	searchParams := make(map[string]interface{})
	if strings.Contains(identity, "@") {
		searchParams["email"] = identity
	} else {
		searchParams["username"] = identity
	}

	Employees, err := db.GetEmployeesWithFilters(searchParams)
	if err != nil {
		return nil, fmt.Errorf("GetEmployee: %v", err)
	}

	var Employee *domain.Employee
	if Employees == nil || len(Employees) <= 0 {
		return nil, nil //fmt.Errorf("GetEmployee: %v", errors.New("Employee with given ID not found!"))
	} else {
		Employee = Employees[0]
	}

	return Employee, nil
}

func (db *dbClient) GetEmployeesWithFilters(searchParams map[string]interface{}) ([]*domain.Employee, error) {
	p := domain.Employee{}

	selectQuery, err := db.GetSelectQueryForStruct(p, searchParams)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesWithFilters: %v", err)
	}

	rows, err := db.RunSelectQuery(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesWithFilters: %v", err)
	}

	Employees, err := GetEmployeesFromRows(rows)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesWithFilters: %v", err)
	}

	return Employees, nil
}

func GetEmployeesFromRows(rows *sql.Rows) ([]*domain.Employee, error) {
	Employees := make([]*domain.Employee, 0)
	for rows.Next() {
		e := domain.Employee{}

		var managerID sql.NullInt64

		err := rows.Scan(&e.EmployeeID, &e.Username, &e.Email, &e.Password, &e.Name, &e.PhoneNumber,
			&e.Address, &e.EmployeeType, &managerID, &e.CreatedAt, &e.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("GetEmployeesFromRows: %v", err)
		}

		if managerID.Valid {
			e.ManagerID = managerID.Int64
		}

		Employees = append(Employees, &e)
	}

	return Employees, nil
}

func (db *dbClient) UpdateEmployee(EmployeeID int64, updates map[string]interface{}) (*domain.Employee, error) {
	updateQuery, err := db.GetUpdateQueryForStruct(domain.Employee{}, EmployeeID, updates)
	if err != nil {
		return nil, fmt.Errorf("UpdateEmployee: %v", err)
	}

	if len(updates) > 0 {
		_, err = db.RunUpdateQuery(updateQuery)
		if err != nil {
			return nil, fmt.Errorf("UpdateEmployee: %v", err)
		}
	}

	Employee, err := db.GetEmployee(EmployeeID)
	if err != nil {
		return nil, fmt.Errorf("UpdateEmployee: %v", err)
	}

	return Employee, nil
}

func (db *dbClient) DeleteEmployee(EmployeeID int64) error {
	deleteParams := make(map[string]interface{})

	deleteParams["employee_id"] = EmployeeID

	deleteQuery, err := db.GetDeleteQueryForStruct(domain.Employee{}, deleteParams)
	if err != nil {
		return fmt.Errorf("DeleteEmployee: %v", err)
	}

	_, err = db.RunDeleteQuery(deleteQuery)
	if err != nil {
		return fmt.Errorf("DeleteEmployee: %v", err)
	}

	return nil
}

func (db *dbClient) CheckForDuplicateEmployee(identity string) (int, error) {
	searchParams := make(map[string]interface{})
	searchKey := ""
	if strings.Contains(identity, "@") {
		searchKey = "email"
	} else {
		searchKey = "username"
	}
	searchParams[searchKey] = identity

	Employees, err := db.GetEmployeesWithFilters(searchParams)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("CheckForDuplicateEmployee: %v", err)
	}
	if len(Employees) > 0 {
		return http.StatusBadRequest,
			fmt.Errorf("CheckForDuplicateEmployee: Employee with this %v already exists", searchKey)
	}

	return http.StatusOK, nil
}

func (db *dbClient) CheckEmployeeLogin(identity, password string) (*domain.Employee, error) {
	Employee, err := db.GetEmployeeWithIdentity(identity)
	if err != nil {
		return nil, err
	}

	if Employee == nil {
		return nil, errors.New("employee with these details does not exist")
	}

	if strings.EqualFold(Employee.Password, password) {
		return Employee, nil
	}

	return nil, errors.New("invalid credentials")
}
