package entity

//- `ID` (int): Unique identifier for the employee.
//
//- `Name` (string): Name of the employee.
//
//- `Position` (string): Position/title of the employee.
//
//- `Salary` (float64): Salary of the employee.

type Employee struct {
	ID       *int     `json:"id"`
	Name     *string  `json:"name"`
	Position *string  `json:"position"`
	Salary   *float64 `json:"salary"`
}
