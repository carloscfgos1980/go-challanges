package datamanagement

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	for i, employee := range m.Employees {
		if employee.ID == id {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
		}
	}
	fmt.Printf("Employee with ID %d removed\n", id)
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	if len(m.Employees) == 0 {
		return 0.0
	}
	total := 0.0
	for _, employee := range m.Employees {
		total += employee.Salary
	}
	return total / float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	for _, employee := range m.Employees {
		if employee.ID == id {
			return &employee
		}
	}
	return nil
}
