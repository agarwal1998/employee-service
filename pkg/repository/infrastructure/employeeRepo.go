package infrastructure

import (
	"context"
	cloudSql "crud/drivers"
	"crud/pkg/entity"
	"fmt"
	"strings"
)

type EmployeeRepo struct {
	sqlClient         *cloudSql.SqlClient
	columns           []string
	EmployeeTableName string
}

func NewEmployeeRepo(sqlClient *cloudSql.SqlClient) *EmployeeRepo {
	return &EmployeeRepo{
		sqlClient,
		[]string{"id", "name", "position", "salary"},
		"employee",
	}
}

func (el *EmployeeRepo) GetEmpoyee(ctx context.Context, employeeId int) (entity.Employee, error) {
	query := fmt.Sprintf("SELECT %v FROM %v where id = ?", strings.Join(el.columns, ","), el.EmployeeTableName)
	stmt, err := el.sqlClient.DB.Prepare(query)
	if err != nil {
		return entity.Employee{}, err
	}
	defer stmt.Close()
	var result entity.Employee
	err = stmt.QueryRow(employeeId).Scan(&result.ID, &result.Name, &result.Position, &result.Salary)
	return result, err
}

func (el *EmployeeRepo) UpdateEmpoyee(ctx context.Context, updateReq entity.Employee) error {
	args := []interface{}{}
	setPart := ""
	if updateReq.Name != nil {
		args = append(args, updateReq.Name)
		setPart += "name=?,"
	}
	if updateReq.Position != nil {
		args = append(args, updateReq.Position)
		setPart += "position=?,"
	}
	if updateReq.Salary != nil {
		args = append(args, updateReq.Salary)
		setPart += "salary=?"
	}
	args = append(args, updateReq.ID)
	query := fmt.Sprintf("UPDATE %v SET %v WHERE id=?", el.EmployeeTableName, setPart)

	stmt, err := el.sqlClient.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}

func (el *EmployeeRepo) DeleteEmpoyee(ctx context.Context, employeeId int) error {
	query := fmt.Sprintf("DELETE FROM %s where id =?", el.EmployeeTableName)
	_, err := el.sqlClient.DB.Exec(query, employeeId)
	return err
}

func (el *EmployeeRepo) Creatempoyee(ctx context.Context, employee entity.Employee) error {
	valueString := strings.Repeat("?,", len(el.columns)-2) + "?"
	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v) ", el.EmployeeTableName, strings.Join(el.columns[1:], ","), valueString)

	stmt, err := el.sqlClient.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(employee.Name, employee.Position, employee.Salary); err != nil {
		return err
	}
	return nil
}
