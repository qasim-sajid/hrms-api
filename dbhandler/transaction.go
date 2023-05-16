package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/qasim-sajid/hrms-api/domain"
)

func (db *dbClient) AddTransaction(transaction *domain.Transaction) (*domain.Transaction, int, error) {
	insertQuery, err := db.GetInsertQuery(*transaction)
	if err != nil {
		return nil, -1, fmt.Errorf("AddTransaction: %v", err)
	}

	_, err = db.RunInsertQuery(insertQuery)
	if err != nil {
		return nil, -1, fmt.Errorf("AddTransaction: %v", err)
	}

	return transaction, http.StatusOK, nil
}

func (db *dbClient) GetAllTransactions() ([]*domain.Transaction, error) {
	transactions, err := db.GetTransactionsWithFilters(make(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("GetAllTransactions: %v", err)
	}

	return transactions, nil
}

func (db *dbClient) GetTransaction(transactionID int64) (*domain.Transaction, error) {
	selectParams := make(map[string]interface{})

	selectParams["transaction_id"] = transactionID

	transactions, err := db.GetTransactionsWithFilters(selectParams)
	if err != nil {
		return nil, fmt.Errorf("GetTransaction: %v", err)
	}

	var transaction *domain.Transaction
	if transactions == nil || len(transactions) <= 0 {
		return nil, fmt.Errorf("GetTransaction: %v", errors.New("transaction with given id not found"))
	} else {
		transaction = transactions[0]
	}

	return transaction, nil
}

func (db *dbClient) GetTransactionsWithFilters(searchParams map[string]interface{}) ([]*domain.Transaction, error) {
	p := domain.Transaction{}

	selectQuery, err := db.GetSelectQueryForStruct(p, searchParams)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionsWithFilters: %v", err)
	}

	rows, err := db.RunSelectQuery(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionsWithFilters: %v", err)
	}

	transactions, err := GetTransactionsFromRows(rows)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionsWithFilters: %v", err)
	}

	return transactions, nil
}

func GetTransactionsFromRows(rows *sql.Rows) ([]*domain.Transaction, error) {
	transactions := make([]*domain.Transaction, 0)
	for rows.Next() {
		c := domain.Transaction{}

		err := rows.Scan(&c.TransactionId, &c.ContractId, &c.TransactionDate, &c.PaidAmount,
			&c.AvailedPto, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("GetTransactionsFromRows: %v", err)
		}

		transactions = append(transactions, &c)
	}

	return transactions, nil
}

func (db *dbClient) UpdateTransaction(transactionID int64, updates map[string]interface{}) (*domain.Transaction, error) {
	updateQuery, err := db.GetUpdateQueryForStruct(domain.Transaction{}, transactionID, updates)
	if err != nil {
		return nil, fmt.Errorf("UpdateTransaction: %v", err)
	}

	if len(updates) > 0 {
		_, err = db.RunUpdateQuery(updateQuery)
		if err != nil {
			return nil, fmt.Errorf("UpdateTransaction: %v", err)
		}
	}

	transaction, err := db.GetTransaction(transactionID)
	if err != nil {
		return nil, fmt.Errorf("UpdateTransaction: %v", err)
	}

	return transaction, nil
}

func (db *dbClient) DeleteTransaction(transactionID int64) error {
	deleteParams := make(map[string]interface{})

	deleteParams["transaction_id"] = transactionID

	deleteQuery, err := db.GetDeleteQueryForStruct(domain.Transaction{}, deleteParams)
	if err != nil {
		return fmt.Errorf("DeleteTransaction: %v", err)
	}

	_, err = db.RunDeleteQuery(deleteQuery)
	if err != nil {
		return fmt.Errorf("DeleteTransaction: %v", err)
	}

	return nil
}
