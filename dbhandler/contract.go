package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/qasim-sajid/hrms-api/models"
)

func (db *dbClient) AddContract(contract *models.Contract) (*models.Contract, int, error) {
	insertQuery, err := db.GetInsertQuery(*contract)
	if err != nil {
		return nil, -1, fmt.Errorf("AddContract: %v", err)
	}

	_, err = db.RunInsertQuery(insertQuery)
	if err != nil {
		return nil, -1, fmt.Errorf("AddContract: %v", err)
	}

	return contract, http.StatusOK, nil
}

func (db *dbClient) GetAllContracts() ([]*models.Contract, error) {
	contracts, err := db.GetContractsWithFilters(make(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("GetAllContracts: %v", err)
	}

	return contracts, nil
}

func (db *dbClient) GetContract(contractID int64) (*models.Contract, error) {
	selectParams := make(map[string]interface{})

	selectParams["contract_id"] = contractID

	contracts, err := db.GetContractsWithFilters(selectParams)
	if err != nil {
		return nil, fmt.Errorf("GetContract: %v", err)
	}

	var contract *models.Contract
	if contracts == nil || len(contracts) <= 0 {
		return nil, fmt.Errorf("GetContract: %v", errors.New("contract with given id not found"))
	} else {
		contract = contracts[0]
	}

	return contract, nil
}

func (db *dbClient) GetContractsWithFilters(searchParams map[string]interface{}) ([]*models.Contract, error) {
	p := models.Contract{}

	selectQuery, err := db.GetSelectQueryForStruct(p, searchParams)
	if err != nil {
		return nil, fmt.Errorf("GetContractsWithFilters: %v", err)
	}

	rows, err := db.RunSelectQuery(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("GetContractsWithFilters: %v", err)
	}

	contracts, err := db.GetContractsFromRows(rows)
	if err != nil {
		return nil, fmt.Errorf("GetContractsWithFilters: %v", err)
	}

	return contracts, nil
}

func (db *dbClient) GetContractsFromRows(rows *sql.Rows) ([]*models.Contract, error) {
	contracts := make([]*models.Contract, 0)
	for rows.Next() {
		c := models.Contract{}

		err := rows.Scan(&c.ContractId, &c.EmployeeID, &c.StartDate, &c.EndDate, &c.BasicPay,
			&c.TotalPto, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("GetContractsFromRows: %v", err)
		}

		contracts = append(contracts, &c)
	}

	return contracts, nil
}

func (db *dbClient) UpdateContract(contractID int64, updates map[string]interface{}) (*models.Contract, error) {
	updateQuery, err := db.GetUpdateQueryForStruct(models.Contract{}, contractID, updates)
	if err != nil {
		return nil, fmt.Errorf("UpdateContract: %v", err)
	}

	if len(updates) > 0 {
		_, err = db.RunUpdateQuery(updateQuery)
		if err != nil {
			return nil, fmt.Errorf("UpdateContract: %v", err)
		}
	}

	contract, err := db.GetContract(contractID)
	if err != nil {
		return nil, fmt.Errorf("UpdateContract: %v", err)
	}

	return contract, nil
}

func (db *dbClient) DeleteContract(contractID int64) error {
	deleteParams := make(map[string]interface{})

	deleteParams["contract_id"] = contractID

	deleteQuery, err := db.GetDeleteQueryForStruct(models.Contract{}, deleteParams)
	if err != nil {
		return fmt.Errorf("DeleteContract: %v", err)
	}

	_, err = db.RunDeleteQuery(deleteQuery)
	if err != nil {
		return fmt.Errorf("DeleteContract: %v", err)
	}

	return nil
}
