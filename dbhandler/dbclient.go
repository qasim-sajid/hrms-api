package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
	"github.com/qasim-sajid/hrms-api/conf"
	"github.com/qasim-sajid/hrms-api/models"
)

type dbClient struct {
	dbName string
}

// NewDBClient returns ref to a new dbClient object
func NewDBClient(dbName string) (h DbHandler, err error) {
	client := &dbClient{
		dbName: dbName,
	}

	return client, nil
}

var dbConnection *sql.DB

// DB set up
func (db *dbClient) SetupDB() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		conf.Configs.DBUser, conf.Configs.DBPassword, conf.Configs.DBName, conf.Configs.DBHost, conf.Configs.DBPort)
	dbC, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(fmt.Errorf("SetupDB: %v", err))
	}

	dbConnection = dbC
}

func (db *dbClient) CloseDB() {
	if dbConnection != nil {
		dbConnection.Close()
	}
}

func (db *dbClient) RunInsertQuery(query string) (sql.Result, error) {
	result, err := dbConnection.Exec(query)

	if err != nil {
		return nil, fmt.Errorf("RunInsertQuery: %v", err)
	}

	return result, nil
}

func (db *dbClient) RunSelectQuery(query string) (*sql.Rows, error) {
	rows, err := dbConnection.Query(query)

	if err != nil {
		return nil, fmt.Errorf("RunSelectQuery: %v", err)
	}

	return rows, nil
}

func (db *dbClient) RunUpdateQuery(query string) (sql.Result, error) {
	result, err := dbConnection.Exec(query)

	if err != nil {
		return nil, fmt.Errorf("RunUpdateQuery: %v", err)
	}

	return result, nil
}

func (db *dbClient) RunDeleteQuery(query string) (sql.Result, error) {
	result, err := dbConnection.Exec(query)

	if err != nil {
		return nil, fmt.Errorf("RunDeleteQuery: %v", err)
	}

	return result, nil
}

func (db *dbClient) GetInsertQuery(structType interface{}) (string, error) {
	tableName, err := db.GetTableNameForStruct(structType)
	if err != nil {
		return ``, fmt.Errorf("GetInsertQuery: %v", err)
	}

	query := ``
	switch reflect.TypeOf(structType).Name() {
	case "Employee":
		employee := structType.(models.Employee)
		query = fmt.Sprintf(`INSERT INTO %s (%s) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%v', '%v'`,
			tableName, db.GetColumnNamesForStruct(employee), employee.Username, employee.Email, employee.Password,
			employee.Name, employee.PhoneNumber, employee.Address, employee.EmployeeType, employee.CreatedAt.Format(conf.TIME_LAYOUT),
			employee.UpdatedAt.Format(conf.TIME_LAYOUT))
		if employee.ManagerID > 0 {
			query = fmt.Sprintf(`%s, %d)`, query, employee.ManagerID)
		} else {
			query = fmt.Sprintf(`%s, %v)`, query, "null")
		}
	case "Contract":
		contract := structType.(models.Contract)
		query = fmt.Sprintf(`INSERT INTO %s (%s) VALUES ('%v', '%v', %f, %f, '%v', '%v', %d)`,
			tableName, db.GetColumnNamesForStruct(contract), contract.StartDate.Format(conf.TIME_LAYOUT),
			contract.EndDate.Format(conf.TIME_LAYOUT), contract.BasicPay, contract.TotalPto,
			contract.CreatedAt.Format(conf.TIME_LAYOUT), contract.UpdatedAt.Format(conf.TIME_LAYOUT), contract.EmployeeID)
	case "Transaction":
		transaction := structType.(models.Transaction)
		query = fmt.Sprintf(`INSERT INTO %s (%s) VALUES ('%v', %f, %f, '%v', '%v', %d)`,
			tableName, db.GetColumnNamesForStruct(transaction), transaction.TransactionDate.Format(conf.TIME_LAYOUT),
			transaction.PaidAmount, transaction.AvailedPto, transaction.CreatedAt.Format(conf.TIME_LAYOUT),
			transaction.UpdatedAt.Format(conf.TIME_LAYOUT), transaction.ContractId)
	case "Request":
		request := structType.(models.Request)
		query = fmt.Sprintf(`INSERT INTO %s (%s) VALUES ('%v', '%v', '%v', %t, '%v', '%v', %d)`,
			tableName, db.GetColumnNamesForStruct(request), request.StartDate.Format(conf.TIME_LAYOUT),
			request.EndDate.Format(conf.TIME_LAYOUT), request.ActionAt.Format(conf.TIME_LAYOUT), request.IsApproved,
			request.CreatedAt.Format(conf.TIME_LAYOUT), request.UpdatedAt.Format(conf.TIME_LAYOUT), request.EmployeeID)
	default:
		return ``, fmt.Errorf("GetInsertQuery: %v",
			errors.New("insert query generation error"))
	}

	return query, nil
}

func (db *dbClient) GetSelectQueryForStruct(structType interface{}, searchParams map[string]interface{}) (string, error) {
	if reflect.ValueOf(structType).Kind() == reflect.Struct {
		tableName, err := db.GetTableNameForStruct(structType)
		if err != nil {
			return ``, fmt.Errorf("GetSelectQueryForStruct: %v", err)
		}

		query := fmt.Sprintf("SELECT * FROM %s", tableName)

		i := 0
		for k, v := range searchParams {
			val := reflect.ValueOf(v)
			kind := val.Kind()
			if kind == reflect.Int || kind == reflect.Int64 {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %d", query, k, val.Int())
				} else {
					query = fmt.Sprintf("%s && %s = %d", query, k, val.Int())
				}
			} else if kind == reflect.Float32 || kind == reflect.Float64 {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %f", query, k, val.Float())
				} else {
					query = fmt.Sprintf("%s && %s = %f", query, k, val.Float())
				}
			} else if kind == reflect.Bool {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %t", query, k, val.Bool())
				} else {
					query = fmt.Sprintf("%s && %s = %t", query, k, val.Bool())
				}
			} else {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = '%s'", query, k, val.String())
				} else {
					query = fmt.Sprintf("%s && %s = '%s'", query, k, val.String())
				}
			}

			i++
		}

		return query, nil
	}

	return ``, fmt.Errorf("GetSelectQueryForStruct: %v", errors.New("select query generation error"))
}

func (db *dbClient) GetUpdateQueryForStruct(structType interface{}, itemID int64, updates map[string]interface{}) (string, error) {
	if reflect.ValueOf(structType).Kind() == reflect.Struct {
		tableName, err := db.GetTableNameForStruct(structType)
		if err != nil {
			return ``, fmt.Errorf("GetUpdateQueryForStruct: %v", err)
		}

		query := fmt.Sprintf("UPDATE %s", tableName)

		i := 0
		for k, v := range updates {
			kind := reflect.ValueOf(v).Kind()
			val := reflect.ValueOf(v)
			if kind == reflect.Int || kind == reflect.Int64 {
				if i == 0 {
					query = fmt.Sprintf("%s SET %s = %d", query, k, val.Int())
				} else {
					query = fmt.Sprintf("%s, %s = %d", query, k, val.Int())
				}
			} else if kind == reflect.Float32 || kind == reflect.Float64 {
				if i == 0 {
					query = fmt.Sprintf("%s SET %s = %f", query, k, val.Float())
				} else {
					query = fmt.Sprintf("%s, %s = %f", query, k, val.Float())
				}
			} else if kind == reflect.Bool {
				if i == 0 {
					query = fmt.Sprintf("%s SET %s = %t", query, k, val.Bool())
				} else {
					query = fmt.Sprintf("%s, %s = %t", query, k, val.Bool())
				}
			} else {
				if i == 0 {
					query = fmt.Sprintf("%s SET %s = '%s'", query, k, val.String())
				} else {
					query = fmt.Sprintf("%s, %s = '%s'", query, k, val.String())
				}
			}

			i++
		}

		idColumn, err := db.GetTableNameForStruct(structType)
		if err != nil {
			return "", err
		}

		query = fmt.Sprintf("%s WHERE %s_id = %d", query, idColumn, itemID)
		return query, nil
	}

	return ``, fmt.Errorf("GetUpdateQueryForStruct: %v", errors.New("update query generation error"))
}

func (db *dbClient) GetDeleteQueryForStruct(structType interface{}, columnParams map[string]interface{}) (string, error) {
	if reflect.ValueOf(structType).Kind() == reflect.Struct {
		tableName, err := db.GetTableNameForStruct(structType)
		if err != nil {
			return ``, fmt.Errorf("GetSelectQueryForStruct: %v", err)
		}

		query := fmt.Sprintf("DELETE FROM %s", tableName)

		i := 0
		for k, v := range columnParams {
			kind := reflect.ValueOf(v).Kind()
			val := reflect.ValueOf(v)
			if kind == reflect.Int || kind == reflect.Int64 {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %d", query, k, val.Int())
				} else {
					query = fmt.Sprintf("%s && %s = %d", query, k, val.Int())
				}
			} else if kind == reflect.Float32 || kind == reflect.Float64 {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %f", query, k, val.Float())
				} else {
					query = fmt.Sprintf("%s && %s = %f", query, k, val.Float())
				}
			} else if kind == reflect.Bool {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = %t", query, k, val.Bool())
				} else {
					query = fmt.Sprintf("%s && %s = %t", query, k, val.Bool())
				}
			} else {
				if i == 0 {
					query = fmt.Sprintf("%s WHERE %s = '%s'", query, k, val.String())
				} else {
					query = fmt.Sprintf("%s && %s = '%s'", query, k, val.String())
				}
			}

			i++
		}

		return query, nil
	}

	return ``, fmt.Errorf("GetSelectQueryForStruct: %v", errors.New("select query generation error"))
}

func (db *dbClient) GetTableNameForStruct(t interface{}) (string, error) {
	switch reflect.TypeOf(t).Name() {
	case "Employee":
		return "employee", nil
	case "Contract":
		return "contract", nil
	case "Transaction":
		return "transactions", nil
	case "Request":
		return "requests", nil
	}

	return "", fmt.Errorf("GetTableNameForStruct: %v", errors.New("Struct not found: "+reflect.TypeOf(t).Name()))
}

func (db *dbClient) GetIDColumnForTable(t interface{}) (string, error) {
	switch reflect.TypeOf(t).Name() {
	case "Employee":
		return "employee_id", nil
	case "Contract":
		return "contract_id", nil
	case "Transaction":
		return "transaction_id", nil
	case "Request":
		return "request_id", nil
	}

	return "", fmt.Errorf("GetIDColumnForTable: %v", errors.New("Struct not found: "+reflect.TypeOf(t).Name()))
}

func (db *dbClient) GetColumnNamesForStruct(structType interface{}) string {
	s := reflect.TypeOf(structType)
	columnNames := ""

	skipCheck := 0

	for i := 0; i < s.NumField(); i++ {
		r := s.Field(i)
		if r.Type.Kind() == reflect.Pointer {
			r = reflect.Indirect(reflect.ValueOf(structType)).Type().Field(i)
		}
		if r.Type.Kind() == reflect.Slice {
			continue
		}
		if r.Type.Kind() == reflect.Array {
			continue
		}

		switch jsonTag := r.Tag.Get("json"); jsonTag {
		case "-":
			continue
		case "":
			continue
		default:
			parts := strings.Split(jsonTag, ",")
			columnName := parts[0]
			if columnName == "" {
				continue
			}

			//Skip the id filed
			idColumn, err := db.GetIDColumnForTable(structType)
			if err != nil {
				panic(fmt.Errorf("GetColumnNamesForStruct: %v", err))
			}
			if columnName == idColumn {
				continue
			}

			if skipCheck == 0 {
				skipCheck++
				columnNames = columnName
			} else {
				columnNames = fmt.Sprintf("%s, %s", columnNames, columnName)
			}
		}
	}

	return columnNames
}
