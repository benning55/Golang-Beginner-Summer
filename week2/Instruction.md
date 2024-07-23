# Week 2: Go Syntax, Data Types, and Control Structures

## Student Work Instructions

### Objective
Build a basic Income Expenses Tracker using Go, focusing on syntax, data types, and control structures.

### Setup
1. Edit file named `main.go`.

### Tasks

1. **Implement Basic Program Structure**
   - From the main function.
   - Implement an infinite loop for the menu system.

2. **Create Variables**
   - Declare variables for `totalIncome` and `totalExpenses` using appropriate data types.

3. **Implement Menu System**
   - Use a `for` loop to continuously display the menu.
   - Implement a `switch` statement to handle user choices.
   - Menu options should include:
     1. Add Income
     2. Add Expense
     3. View Summary
     4. Exit

4. **Add Income Functionality**
   - Prompt user for income amount.
   - Validate input (ensure it's a non-negative number).
   - Add valid income to `totalIncome`.
   - Display confirmation message.

5. **Add Expense Functionality**
   - Prompt user for expense amount.
   - Validate input (ensure it's a non-negative number).
   - Add valid expense to `totalExpenses`.
   - Display confirmation message.

6. **Implement Summary View**
   - Display total income.
   - Display total expenses.
   - Calculate and display net savings (income - expenses).

7. **Implement Exit Option**
   - Use `os.Exit(0)` to terminate the program when user chooses to exit.

8. **Error Handling**
   - Implement a default case in the switch statement to handle invalid menu choices.

### Example Code Structure

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Welcome message
	// Declare variables

	for {
		// Display menu
		// Get user choice

		switch choice {
		case 1:
			// Add Income
		case 2:
			// Add Expense
		case 3:
			// View Summary
		case 4:
			// Exit
		default:
			// Handle invalid choice
		}
	}
}