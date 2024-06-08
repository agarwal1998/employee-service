package logic

import (
	"context"
	"crud/common/helper"
	"crud/mocks"
	"crud/pkg/entity"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEmpoyee(t *testing.T) {
	employeeRepoMock := new(mocks.EmployeeRepoMock)
	employeeLogic := NewEmployeeLogic(employeeRepoMock)
	employee := entity.Employee{
		helper.GetIntPointer(1), helper.GetStringPointer("aayush"),
		helper.GetStringPointer("SDE-2"), helper.GetFloat64Pointer(10),
	}

	employeeRepoMock.On("GetEmpoyee", mock.Anything, mock.Anything).Return(entity.Employee{
		helper.GetIntPointer(1), helper.GetStringPointer("aayush"),
		helper.GetStringPointer("SDE-2"), helper.GetFloat64Pointer(10),
	}, nil)

	employeeRepoMock.On("Creatempoyee", mock.Anything, mock.Anything).Return(nil)

	employeeRepoMock.On("DeleteEmpoyee", mock.Anything, mock.Anything).Return(nil)

	employeeRepoMock.On("UpdateEmpoyee", mock.Anything, mock.Anything).Return(nil)

	t.Run("get employee", func(t *testing.T) {
		resp, err := employeeLogic.GetEmpoyee(context.Background(), 1)
		assert.Equal(t, resp, employee)
		assert.Equal(t, err, nil)
	})
	t.Run("create employee", func(t *testing.T) {
		err := employeeLogic.Creatempoyee(context.Background(), employee)
		assert.Equal(t, err, nil)
	})
	t.Run("delete employee", func(t *testing.T) {
		err := employeeLogic.DeleteEmpoyee(context.Background(), 1)
		assert.Equal(t, err, nil)
	})
	t.Run("update employee", func(t *testing.T) {
		err := employeeLogic.UpdateEmpoyee(context.Background(), employee)
		assert.Equal(t, err, nil)
	})

}
