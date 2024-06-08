package repository

import (
	"context"
	"crud/pkg/entity"
)

type UserRepoInterface interface {
	GetEmpoyee(ctx context.Context, employeeId int) (entity.Employee, error)
	UpdateEmpoyee(ctx context.Context, updateReq entity.Employee) error
	DeleteEmpoyee(ctx context.Context, employeeId int) error
	Creatempoyee(ctx context.Context, employee entity.Employee) error
}
