package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	datamanagement "github.com/carloscfgos1980/go-challanges/7.data_management"
	reversestring "github.com/carloscfgos1980/go-challanges/8.reverse_string"
)

func main() {

	commands := map[string]func(){
		"data": func() {
			manager := datamanagement.Manager{}
			manager.AddEmployee(datamanagement.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
			manager.AddEmployee(datamanagement.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
			manager.AddEmployee(datamanagement.Employee{ID: 3, Name: "Charlie", Age: 28, Salary: 72000})
			fmt.Printf("Initial Employees: %+v\n", manager.Employees)
			manager.RemoveEmployee(2)
			fmt.Printf("Employees after removal: %+v\n", manager.Employees)
			averageSalary := manager.GetAverageSalary()
			fmt.Printf("Average Salary: %f\n", averageSalary)
			employee := manager.FindEmployeeByID(1)
			if employee != nil {
				fmt.Printf("Employee with ID 1: %+v\n", *employee)
			} else {
				fmt.Printf("Employee with ID 1 not found\n")
			}
		},
		"reverse": func() {
			if len(os.Args) < 3 {
				fmt.Printf("usage: go run . reverse <string>\n")
				return
			}
			input := os.Args[2]
			output := reversestring.ReverseString(input)
			fmt.Println(output)
		},
	}
	if len(os.Args) < 2 {
		fmt.Printf("usage: go run . [%s]\n", strings.Join(commandNames(commands), "|"))
		return
	}

	run, ok := commands[os.Args[1]]
	if !ok {
		fmt.Printf("unknown command %q. available: %s\n", os.Args[1], strings.Join(commandNames(commands), ", "))
		return
	}

	run()

}

func commandNames(commands map[string]func()) []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
