# Week 3: Functions and Collections in Go

## Student Work Instructions

### Objective
Enhance the Income Expenses Tracker using Go, focusing on functions, slices, and maps.

### Setup
1. You will have two files named `main.go` and `operations.go`.

### Tasks

1. **Implement Basic Program Structure**
   - Use the provided `main` function in `main.go`.
   - Implement an infinite loop for the menu system.

2. **Create Variables**
   - Use the provided global variables: `TotalIncome`, `TotalExpenses`, `Transactions`, and `ExpenseCategories`.

3. **Implement Menu System**
   - Use the provided `displayMenu` function.
   - Implement a `switch` statement to handle user choices:
     1. Add Income
     2. Add Expense
     3. View Summary
     4. View Transactions
     5. View Expenses By Category
     6. Exit

4. **Add Income Functionality**
   - Implement the `AddIncome` function in `operations.go`.
   - Prompt user for income amount and validate input.
   - Update `TotalIncome` and append to `Transactions`.

5. **Add Expense Functionality**
   - Implement the `AddExpense` function in `operations.go`.
   - Prompt user for expense amount and validate input.
   - Use `getCategory` function to get the expense category.
   - Update `TotalExpenses`, `Transactions`, and `ExpenseCategories`.

6. **Implement Summary View**
   - Implement the `ViewSummary` function to display total income, total expenses, and net savings.

7. **Implement Transaction View**
   - Implement the `ViewTransactions` function to display all transactions with their type (income or expense).

8. **Implement Expenses By Category View**
   - Implement the `ViewExpensesByCategory` function to display expenses grouped by category.

9. **Implement Exit Option**
   - Use the provided `ExitProgram` function to terminate the program.

### Key Concepts to Focus On

1. **Functions**: 
   - Implement and use functions for different operations.
   - Understand function parameters and return values.

2. **Slices**: 
   - Use the `Transactions` slice to store transaction history.
   - Implement operations to append to and iterate over slices.

3. **Maps**: 
   - Use the `ExpenseCategories` map to categorize expenses.
   - Implement operations to add to and display data from maps.

4. **Basic Input Validation**: 
   - Ensure amounts entered are non-negative.

5. **Control Structures**: 
   - Use `for` loops, `if` statements, and `switch` statements effectively.

### Hints

- Use `fmt.Print` and `fmt.Scan` for user input and output.
- Remember to update both `Transactions` and `ExpenseCategories` when adding an expense.
- When displaying transactions, use conditional statements to differentiate between income and expenses.

### Running the Program

1. Open a terminal and navigate to your project directory.
2. Run the program using:
   ```sh
   go run .