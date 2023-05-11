package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/qasim-sajid/hrms-api/models"
)

func (db *dbClient) AddRequest(request *models.Request) (*models.Request, int, error) {
	insertQuery, err := db.GetInsertQuery(*request)
	if err != nil {
		return nil, -1, fmt.Errorf("AddRequest: %v", err)
	}

	_, err = db.RunInsertQuery(insertQuery)
	if err != nil {
		return nil, -1, fmt.Errorf("AddRequest: %v", err)
	}

	return request, http.StatusOK, nil
}

func (db *dbClient) GetAllRequests() ([]*models.Request, error) {
	requests, err := db.GetRequestsWithFilters(make(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("GetAllRequests: %v", err)
	}

	return requests, nil
}

func (db *dbClient) GetRequest(requestID int64) (*models.Request, error) {
	selectParams := make(map[string]interface{})

	selectParams["request_id"] = requestID

	requests, err := db.GetRequestsWithFilters(selectParams)
	if err != nil {
		return nil, fmt.Errorf("GetRequest: %v", err)
	}

	var request *models.Request
	if requests == nil || len(requests) <= 0 {
		return nil, fmt.Errorf("GetRequest: %v", errors.New("request with given id not found"))
	} else {
		request = requests[0]
	}

	return request, nil
}

func (db *dbClient) GetRequestsWithFilters(searchParams map[string]interface{}) ([]*models.Request, error) {
	p := models.Request{}

	selectQuery, err := db.GetSelectQueryForStruct(p, searchParams)
	if err != nil {
		return nil, fmt.Errorf("GetRequestsWithFilters: %v", err)
	}

	rows, err := db.RunSelectQuery(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("GetRequestsWithFilters: %v", err)
	}

	requests, err := db.GetRequestsFromRows(rows)
	if err != nil {
		return nil, fmt.Errorf("GetRequestsWithFilters: %v", err)
	}

	return requests, nil
}

func (db *dbClient) GetRequestsFromRows(rows *sql.Rows) ([]*models.Request, error) {
	requests := make([]*models.Request, 0)
	for rows.Next() {
		c := models.Request{}

		err := rows.Scan(&c.RequestId, &c.EmployeeID, &c.StartDate, &c.EndDate, &c.ActionAt,
			&c.IsApproved, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("GetRequestsFromRows: %v", err)
		}

		requests = append(requests, &c)
	}

	return requests, nil
}

func (db *dbClient) UpdateRequest(requestID int64, updates map[string]interface{}) (*models.Request, error) {
	updateQuery, err := db.GetUpdateQueryForStruct(models.Request{}, requestID, updates)
	if err != nil {
		return nil, fmt.Errorf("UpdateRequest: %v", err)
	}

	if len(updates) > 0 {
		_, err = db.RunUpdateQuery(updateQuery)
		if err != nil {
			return nil, fmt.Errorf("UpdateRequest: %v", err)
		}
	}

	request, err := db.GetRequest(requestID)
	if err != nil {
		return nil, fmt.Errorf("UpdateRequest: %v", err)
	}

	return request, nil
}

func (db *dbClient) DeleteRequest(requestID int64) error {
	deleteParams := make(map[string]interface{})

	deleteParams["request_id"] = requestID

	deleteQuery, err := db.GetDeleteQueryForStruct(models.Request{}, deleteParams)
	if err != nil {
		return fmt.Errorf("DeleteRequest: %v", err)
	}

	_, err = db.RunDeleteQuery(deleteQuery)
	if err != nil {
		return fmt.Errorf("DeleteRequest: %v", err)
	}

	return nil
}
