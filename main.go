package main

import (
	"fmt"

	"github.com/qasim-sajid/hrms-api/conf"
	"github.com/qasim-sajid/hrms-api/dbhandler"
)

// Main function
func main() {
	conf.InitConfigs()

	dbC, err := dbhandler.NewDBClient(conf.Configs.DBName)
	if err != nil {
		panic(fmt.Errorf("NewHandler: %v", err))
	}

	//Setup DB Connection
	dbC.SetupDB()
	defer dbC.CloseDB()

	// employee := models.Employee{}
	// employee.Username = "employee2"
	// employee.Email = "employee2@gmail.com"
	// employee.Password = "employee2pass"
	// employee.Name = "Employee 2"
	// employee.PhoneNumber = "03001234567"
	// employee.Address = "Employee 2 house"
	// employee.EmployeeType = "Contractor"
	// employee.ManagerID = 1
	// employee.CreatedAt = time.Now()
	// employee.UpdatedAt = time.Now()

	// _, status, err := dbC.AddEmployee(&employee)
	// if err != nil {
	// 	panic(err)
	// }

	// contract := models.Contract{}
	// contract.StartDate = time.Now()
	// contract.EndDate = time.Now()
	// contract.BasicPay = 55000
	// contract.TotalPto = 240
	// contract.EmployeeID = 1
	// contract.CreatedAt = time.Now()
	// contract.UpdatedAt = time.Now()

	// _, status, err := dbC.AddContract(&contract)
	// if err != nil {
	// 	panic(err)
	// }

	// transaction := models.Transaction{}
	// transaction.TransactionDate = time.Now()
	// transaction.PaidAmount = 4500
	// transaction.AvailedPto = 20.5
	// transaction.CreatedAt = time.Now()
	// transaction.UpdatedAt = time.Now()
	// transaction.ContractId = 1

	// _, status, err := dbC.AddTransaction(&transaction)
	// if err != nil {
	// 	panic(err)
	// }

	// request := models.Request{}
	// request.StartDate = time.Now()
	// request.EndDate = time.Now()
	// request.ActionAt = time.Now()
	// request.IsApproved = false
	// request.EmployeeID = 1
	// request.CreatedAt = time.Now()
	// request.UpdatedAt = time.Now()

	// _, status, err := dbC.AddRequest(&request)
	// if err != nil {
	// 	panic(err)
	// }

	employees, err := dbC.GetAllEmployees()
	if err != nil {
		panic(err)
	}

	for _, e := range employees {
		fmt.Println("Employees: ", e.Name)
	}

	// fmt.Println("Status: ", status)
}
